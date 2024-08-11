module GraphqlSchema
  module Objects
    module Products
      module IosysCrawlSettings
        module CreateIosysCrawlSettingExcludeKeyword
          class CreateIosysCrawlSettingExcludeKeywordResultValidationFailed < Base
            implements Interfaces::UserError

            field :details, [ErrorDetailType], null: false
          end
        end
      end
    end
  end
end
