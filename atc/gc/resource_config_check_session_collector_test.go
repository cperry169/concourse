package gc_test

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/gc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceConfigCheckSessionCollector", func() {
	var (
		collector                           gc.Collector
		resourceConfigCheckSessionLifecycle db.ResourceConfigCheckSessionLifecycle
		resourceConfigCheckSessionFactory   db.ResourceConfigCheckSessionFactory
		resourceConfigCheckSession          db.ResourceConfigCheckSession
		ownerExpiries                       db.ContainerOwnerExpiries
		resource                            db.Resource
	)

	BeforeEach(func() {
		resourceConfigCheckSessionLifecycle = db.NewResourceConfigCheckSessionLifecycle(dbConn)
		resourceConfigCheckSessionFactory = db.NewResourceConfigCheckSessionFactory(dbConn, lockFactory)
		collector = gc.NewResourceConfigCheckSessionCollector(resourceConfigCheckSessionLifecycle)
	})

	Describe("Run", func() {
		var runErr error

		ownerExpiries = db.ContainerOwnerExpiries{
			GraceTime: 5 * time.Second,
			Min:       10 * time.Second,
			Max:       10 * time.Second,
		}

		BeforeEach(func() {
			var err error
			var found bool
			resource, found, err = defaultPipeline.Resource("some-resource")
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())

			resourceConfigCheckSession, err = resourceConfigCheckSessionFactory.FindOrCreateResourceConfigCheckSession(logger,
				"some-base-type",
				atc.Source{
					"some": "source",
				},
				creds.VersionedResourceTypes{},
				ownerExpiries,
			)
			Expect(err).ToNot(HaveOccurred())

			err = resource.SetResourceConfig(resourceConfigCheckSession.ResourceConfig().ID())
			Expect(err).ToNot(HaveOccurred())
		})

		resourceConfigCheckSessionExists := func(resourceConfigCheckSession db.ResourceConfigCheckSession) bool {
			var count int
			err = psql.Select("COUNT(*)").
				From("resource_config_check_sessions").
				Where(sq.Eq{"id": resourceConfigCheckSession.ID()}).
				RunWith(dbConn).
				QueryRow().
				Scan(&count)
			Expect(err).NotTo(HaveOccurred())
			return count == 1
		}

		JustBeforeEach(func() {
			runErr = collector.Run(context.TODO())
			Expect(runErr).ToNot(HaveOccurred())
		})

		Context("when the resource config check session is expired", func() {
			BeforeEach(func() {
				time.Sleep(ownerExpiries.Max)
			})

			It("removes the resource config check session", func() {
				Expect(resourceConfigCheckSessionExists(resourceConfigCheckSession)).To(BeFalse())
			})
		})

		Context("when the resource config check session is not expired", func() {
			It("keeps the resource config check session", func() {
				Expect(resourceConfigCheckSessionExists(resourceConfigCheckSession)).To(BeTrue())
			})
		})

		Context("when the resource is active and unpaused", func() {
			It("keeps the resource config check session", func() {
				Expect(resourceConfigCheckSessionExists(resourceConfigCheckSession)).To(BeTrue())
			})
		})

		Context("when the resource is unactive", func() {
			BeforeEach(func() {
				atcConfig := atc.Config{
					Jobs: atc.JobConfigs{
						{
							Name: "some-job",
						},
						{
							Name: "some-other-job",
						},
					},
				}

				defaultPipeline, _, err = defaultTeam.SavePipeline("default-pipeline", atcConfig, db.ConfigVersion(1), db.PipelineUnpaused)
				Expect(err).NotTo(HaveOccurred())
			})

			It("removes the resource config check session", func() {
				Expect(resourceConfigCheckSessionExists(resourceConfigCheckSession)).To(BeFalse())
			})
		})

		Context("when the resource is paused", func() {
			BeforeEach(func() {
				var err error
				err = resource.Pause()
				Expect(err).ToNot(HaveOccurred())
			})

			It("removes the resource config check session", func() {
				Expect(resourceConfigCheckSessionExists(resourceConfigCheckSession)).To(BeFalse())
			})
		})
	})
})
