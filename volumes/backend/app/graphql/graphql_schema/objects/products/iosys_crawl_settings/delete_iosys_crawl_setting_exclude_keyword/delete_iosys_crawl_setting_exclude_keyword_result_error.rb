module GraphqlSchema
  module Objects
    module Products
      module IosysCrawlSettings
        module DeleteIosysCrawlSettingExcludeKeyword
          class DeleteIosysCrawlSettingExcludeKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::IosysCrawlSettings::DeleteIosysCrawlSettingExcludeKeyword::
                  DeleteIosysCrawlSettingExcludeKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
