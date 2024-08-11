module GraphqlSchema
  module Objects
    module Products
      module PcKoubouCrawlSettings
        module CreatePcKoubouCrawlSettingExcludeKeyword
          class CreatePcKoubouCrawlSettingExcludeKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::PcKoubouCrawlSettings::CreatePcKoubouCrawlSettingExcludeKeyword::
                  CreatePcKoubouCrawlSettingExcludeKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
