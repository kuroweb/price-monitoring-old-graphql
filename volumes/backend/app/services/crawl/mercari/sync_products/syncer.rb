module Crawl
  module Mercari
    module SyncProducts
      class Syncer
        def self.call(...)
          new(...).call
        end

        def initialize(product:)
          @product = product
        end

        def call
          return unless mercari_crawl_setting.enabled?

          crawl_results = fetch_crawl_results
          save(crawl_results)
        end

        private

        attr_reader :product

        def fetch_crawl_results
          crawl_results = Crawler.new(product:).execute
          InspectCrawlResults.new(mercari_crawl_setting:, crawl_results:).execute
        end

        def save(crawl_results)
          MercariProduct.transaction do
            upsert(crawl_results)
            enqueue_part_sync(crawl_results)
          end
        end

        def upsert(crawl_results)
          upsert_params = crawl_results.results.map do |result|
            result.attributes.merge("product_id" => product.id)
          end

          MercariProduct.upsert_all(upsert_params, record_timestamps: true)
        end

        def enqueue_part_sync(crawl_results)
          not_exist_mercari_products =
            MercariProduct
            .where(product_id: product.id, published: true)
            .where.not(mercari_id: crawl_results.results.map(&:mercari_id))

          job_params = not_exist_mercari_products.map { |mercari_product| [mercari_product.id] }
          Crawl::Mercari::SyncProduct::SyncJob.perform_bulk(job_params)
        end

        def mercari_crawl_setting
          @mercari_crawl_setting ||= product.mercari_crawl_setting
        end
      end
    end
  end
end
