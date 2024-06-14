module GraphqlSchema
  module Objects
    module Products
      module MercariCrawlSettings
        module UpdateMercariCrawlSettingExcludeKeyword
          class UpdateMercariCrawlSettingExcludeKeywordResultSuccess < Base
            implements Interfaces::Products::ResultBase

            field :mercari_crawl_setting_exclude_keyword,
                  Objects::Products::MercariCrawlSettings::MercariCrawlSettingExcludeKeyword, null: false
          end
        end
      end
    end
  end
end
