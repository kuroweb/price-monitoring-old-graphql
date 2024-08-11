module GraphqlSchema
  module Objects
    module Products
      module IosysCrawlSettings
        module DeleteIosysCrawlSettingRequiredKeyword
          class DeleteIosysCrawlSettingRequiredKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::IosysCrawlSettings::DeleteIosysCrawlSettingRequiredKeyword::
                  DeleteIosysCrawlSettingRequiredKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
