package translator

import (
	"sync"
)

type Language string

const (
	EN Language = "en"
	RU Language = "ru"
)

// Translation keys constants
const (
	// Universal fields
	KeyTitleRequired               = "title_required"
	KeyDescriptionRequired         = "description_required"
	KeyLocationRequired            = "location_required"
	KeyCategoryRequired            = "category_required"
	KeyPriceRequired               = "price_required"
	KeySalaryRequired              = "salary_required"
	KeyContactInfoRequired         = "contact_info_required"
	KeyPropertyTypeRequired        = "property_type_required"
	KeyRoomCountRequired           = "room_count_required"
	KeyFloorRequired               = "floor_required"
	KeyTotalFloorsRequired         = "total_floors_required"
	KeyAreaRequired                = "area_required"
	KeyExperienceRequired          = "experience_required"
	KeyEducationRequired           = "education_required"
	KeySkillsRequired              = "skills_required"
	KeyCoverLetterRequired         = "cover_letter_required"
	KeyMessageRequired             = "message_required"
	KeyNotificationTypeRequired    = "notification_type_required"
	KeyNotificationEnabledRequired = "notification_enabled_required"
	KeySearchQueryRequired         = "search_query_required"

	// Universal IDs
	KeyIDRequired             = "id_required"
	KeyUserIDRequired         = "user_id_required"
	KeyGroupIDRequired        = "group_id_required"
	KeyLocationIDRequired     = "location_id_required"
	KeySiteCategoryIDRequired = "site_category_id_required"

	// Universal operations
	KeyItemCreatedSuccess = "item_created_success"
	KeyItemUpdatedSuccess = "item_updated_success"
	KeyItemDeletedSuccess = "item_deleted_success"
	KeyItemNotFound       = "item_not_found"
	KeyItemCreationError  = "item_creation_error"
	KeyItemUpdateError    = "item_update_error"
	KeyItemDeleteError    = "item_delete_error"
	KeyItemCreationFailed = "item_creation_failed"
	KeyItemUpdateFailed   = "item_update_failed"
	KeyItemDeleteFailed   = "item_delete_failed"

	// Authentication
	KeyRegistrationError                 = "registration_error"
	KeyRegistrationCompletedSuccessfully = "registration_completed_successfully"
	KeyUserNotFound                      = "user_not_found"
	KeyEmailNotConfirmed                 = "email_not_confirmed"
	KeyInvalidLoginOrPassword            = "invalid_login_or_password"
	KeyTokenGenerationError              = "token_generation_error"
	KeyRefreshTokenInvalid               = "refresh_token_invalid"
	KeyRefreshTokenConvertError          = "refresh_token_convert_error"
	KeyJWTTokenCreationError             = "jwt_token_creation_error"
	KeyUserAlreadyExists                 = "user_already_exists"
	KeyInvalidSigningMethod              = "invalid_signing_method"
	KeyTokenClaimsError                  = "token_claims_error"

	// Дополнительные поля аутентификации
	KeyUsernameRequired        = "username_required"
	KeyEmailRequired           = "email_required"
	KeyPasswordRequired        = "password_required"
	KeyPasswordConfirmRequired = "password_confirm_required"
	KeyPhoneRequired           = "phone_required"
	KeyResumeRequired          = "resume_required"

	// Дополнительные ошибки и сообщения
	KeyUserCreationError                      = "user_creation_error"
	KeyPasswordConfirmMismatch                = "password_confirm_mismatch"
	KeyDatabaseError                          = "database_error"
	KeyPasswordRestoreError                   = "password_restore_error"
	KeyUserRolesFetchError                    = "user_roles_fetch_error"
	KeyNoData                                 = "no_data"
	KeyFailedToGenerateID                     = "failed_to_generate_id"
	KeyPhoneImageGenerationError              = "phone_image_generation_error"
	KeyErrorSendingToMessengers               = "error_sending_to_messengers"
	KeyUserIDShouldNotBeNil                   = "user_id_should_not_be_nil"
	KeyIDShouldNotBeNil                       = "id_should_not_be_nil"
	KeyOrdersFetchError                       = "orders_fetch_error"
	KeySerializationError                     = "serialization_error"
	KeyOrderUpdateError                       = "order_update_error"
	KeyDataNotFound                           = "data_not_found"
	KeyResumeIDShouldNotBeNil                 = "resume_id_should_not_be_nil"
	KeyGroupIDInvalidFormat                   = "group_id_invalid_format"
	KeyMarshalError                           = "marshal_error"
	KeyInputDataIncorrectFormat               = "input_data_incorrect_format"
	KeyUsernameAlphanum                       = "username_alphanum"
	KeyUsernameMustConsistOfLettersAndNumbers = "username_must_consist_of_letters_and_numbers"
	KeyEmailInvalid                           = "email_invalid"
	KeyPasswordMinLength                      = "password_min_length"
	KeyPasswordMaxLength                      = "password_max_length"

	// User related
	KeyUserIDEmpty         = "user_id_empty"
	KeyUserRoleNotFound    = "user_role_not_found"
	KeyUserRoleUpdateError = "user_role_update_error"

	// Vacancy related
	KeyVacancyIDShouldNotBeNil = "vacancy_id_should_not_be_nil"

	// Resume related
	KeyResumeNotFound = "resume_not_found"

	// Messenger related
	KeyMessengerGetError = "messenger_get_error"

	// Skill related
	KeySkillNameRequired     = "skill_name_required"
	KeyInsertError           = "insert_error"
	KeySkillDeleteIDRequired = "skill_delete_id_required"

	// Currency related
	KeyCurrencyNotFound = "currency_not_found"

	// Vacancy respond related
	KeyVacancyRespondIDRequired = "vacancy_respond_id_required"

	// User order related
	KeyUserOrderIDRequired = "user_order_id_required"

	// Other
	KeyFilterBuildError     = "filter_build_error"
	KeyDataProcessingError  = "data_processing_error"
	KeyDatabaseDataError    = "database_data_error"
	KeyNameRequired         = "name_required"
	KeyInvalidMetaParameter = "invalid_meta_parameter"

	// Дополнительные константы для всех ключей в карте переводов
	KeyPasswordHashError                  = "password_hash_error"
	KeyActivationCodeInvalid              = "activation_code_invalid"
	KeyEmailConfirmationSuccess           = "email_confirmation_success"
	KeyPasswordRestoreSuccess             = "password_restore_success"
	KeyUserRoleUpdateSuccess              = "user_role_update_success"
	KeyUserHasRoleError                   = "user_has_role_error"
	KeyAuthorizationCompletedSuccessfully = "authorization_completed_successfully"
	KeyUserIdentificationError            = "user_identification_error"
	KeyUserRolesAndPermissions            = "user_roles_and_permissions"
	KeyInvalidToken                       = "invalid_token"
	KeyLoginError                         = "login_error"
	KeyKeySuccessfullyObtained            = "key_successfully_obtained"
	KeyAccountConfirmationError           = "account_confirmation_error"
	KeyInvalidEmailAddress                = "invalid_email_address"
	KeyFailedToRequestPasswordReset       = "failed_to_request_password_reset"
	KeyPasswordUpdated                    = "password_updated"
	KeyValidationError                    = "validation_error"
	KeyUnauthorized                       = "unauthorized"
	KeyForbidden                          = "forbidden"
	KeyNotFound                           = "not_found"
	KeyInternalServerError                = "internal_server_error"
	KeyServiceUnavailable                 = "service_unavailable"
	KeyMissingRequiredParams              = "missing_required_params"
	KeyDataCouldNotBeReceived             = "data_could_not_be_received"
	KeyData                               = "data"
	KeyInvalidLimitParam                  = "invalid_limit_param"
	KeySuccess                            = "success"
	KeyRecordUpdatedSuccessfully          = "record_updated_successfully"
	KeyErrorDeletingData                  = "error_deleting_data"
	KeyFillUserID                         = "fill_user_id"
	KeyNotEnoughRights                    = "not_enough_rights"
	KeyUserUpdateError                    = "user_update_error"
	KeyErrorProcessingForm                = "error_processing_form"
	KeyDataMissingInDataField             = "data_missing_in_data_field"
	KeyIncorrectDataInDataField           = "incorrect_data_in_data_field"
	KeyInvalidInput                       = "invalid_input"
	KeyNoDataForCreation                  = "no_data_for_creation"
	KeyNoAccess                           = "no_access"
	KeyIncorrectJSONData                  = "incorrect_json_data"
	KeyResumes                            = "resumes"
	KeyResume                             = "resume"
	KeyResumeDoesNotBelongToUser          = "resume_does_not_belong_to_user"
	KeyNoActiveResume                     = "no_active_resume"
	KeyVacancies                          = "vacancies"
	KeyVacancy                            = "vacancy"
	KeyVacancySuccessfullyAdded           = "vacancy_successfully_added"
	KeyPhone                              = "phone"
	KeyNoPhone                            = "no_phone"
	KeyFailedToRespond                    = "failed_to_respond"
	KeyRespondedSuccessfully              = "responded_successfully"
	KeyVacancyResponds                    = "vacancy_responds"
	KeyVacancyRespond                     = "vacancy_respond"
	KeyLocationError                      = "location_error"
	KeyQueryTooShort                      = "query_too_short"
	KeySearchResult                       = "search_result"

	// Дополнительные константы для строковых ключей
	KeyNoFilesInForm                            = "no_files_in_form"
	KeyDirectoryCreationError                   = "directory_creation_error"
	KeyDiskSpaceInsufficient                    = "disk_space_insufficient"
	KeyFileOpenError                            = "file_open_error"
	KeyFileDeletionError                        = "file_deletion_error"
	KeyDiskSpaceCheckError                      = "disk_space_check_error"
	KeyDiskSpaceInsufficientDetailed            = "disk_space_insufficient_detailed"
	KeyFilesystemStatsError                     = "filesystem_stats_error"
	KeyInvalidPathError                         = "invalid_path_error"
	KeyDiskInfoError                            = "disk_info_error"
	KeyLocationGetError                         = "location_get_error"
	KeySiteCategoryGetError                     = "site_category_get_error"
	KeyQuickAnnouncementIDShouldNotBeNil        = "quick_announcement_id_should_not_be_nil"
	KeyQuickAnnouncementRespondIDShouldNotBeNil = "quick_announcement_respond_id_should_not_be_nil"
	KeyQuickAnnouncementRequestIsNil            = "quick_announcement_request_is_nil"
	KeyQuickAnnouncementRespondRequestIsNil     = "quick_announcement_respond_request_is_nil"
	KeyUUIDGenerationError                      = "uuid_generation_error"
	KeyTxIsNil                                  = "tx_is_nil"
	KeyFilterQueryIsNil                         = "filter_query_is_nil"
	KeyIDParameterIsRequired                    = "id_parameter_is_required"
	KeyRealty                                   = "realty"
	KeyUserRoleSetFailed                        = "user_role_set_failed"
	KeyPasswordUpdateError                      = "password_update_error"
	KeyTransactionCommitError                   = "transaction_commit_error"
	KeyUserSearchError                          = "user_search_error"
	KeyUserRolesPermissionsError                = "user_roles_permissions_error"
	KeyItemTypeIDRequired                       = "item_type_id_required"
	KeyUsernameAlreadyRegistered                = "username_already_registered"
	KeyCategoryNameError                        = "category_name_error"
	KeyUserRegistrationFailed                   = "user_registration_failed"
	KeyActivationCodeNotFound                   = "activation_code_not_found"
	KeyOfferingIDRequired                       = "offering_id_required"
	KeyEmailConfirmationError                   = "email_confirmation_error"
	KeyEmailValidationError                     = "email_validation_error"
	KeyTransactionCreationError                 = "transaction_creation_error"
	KeyUserFilterError                          = "user_filter_error"
	KeyEmailAlreadyRegistered                   = "email_already_registered"
	KeyInvalidActivationCode                    = "invalid_activation_code"
	KeyOfferings                                = "offerings"
	KeyLocations                                = "locations"
	KeySiteCategories                           = "site_categories"
	KeySettings                                 = "settings"
)

var (
	translations = map[string]map[Language]string{
		// Основные ошибки валидации
		KeyInputDataIncorrectFormat: {
			EN: "Input data is in an incorrect format",
			RU: "Данные введены в неверном формате",
		},
		KeyUsernameRequired: {
			EN: "Username is required",
			RU: "Имя пользователя обязательно",
		},
		KeyUsernameAlphanum: {
			EN: "Username must consist of letters and numbers",
			RU: "Имя пользователя должно состоять из английских букв и цифр",
		},
		KeyUsernameMustConsistOfLettersAndNumbers: {
			EN: "Username must consist of letters and numbers",
			RU: "Имя пользователя должно состоять из букв и цифр",
		},
		KeyEmailRequired: {
			EN: "Email is required",
			RU: "Email обязателен",
		},
		KeyEmailInvalid: {
			EN: "Invalid email format",
			RU: "Неверный формат email",
		},
		KeyPasswordRequired: {
			EN: "Password is required",
			RU: "Пароль обязателен",
		},
		KeyPasswordMinLength: {
			EN: "Password must be at least 8 characters",
			RU: "Пароль должен быть не менее 8 символов",
		},
		KeyPasswordMaxLength: {
			EN: "Password must be no more than 32 characters",
			RU: "Пароль должен быть не более 32 символов",
		},
		KeyPasswordConfirmRequired: {
			EN: "Password confirmation is required",
			RU: "Подтверждение пароля обязательно",
		},
		KeyPasswordConfirmMismatch: {
			EN: "Password confirmation does not match",
			RU: "Подтверждение пароля не совпадает",
		},
		KeyPhoneRequired: {
			EN: "Phone is required",
			RU: "Телефон обязателен",
		},

		// Универсальные поля для всех сущностей
		KeyTitleRequired: {
			EN: "Title is required",
			RU: "Название обязательно",
		},
		KeyDescriptionRequired: {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		KeyLocationRequired: {
			EN: "Location is required",
			RU: "Местоположение обязательно",
		},
		KeyLocationIDRequired: {
			EN: "Location ID is required",
			RU: "ID местоположения обязателен",
		},
		KeyCategoryRequired: {
			EN: "Category is required",
			RU: "Категория обязательна",
		},
		KeySiteCategoryIDRequired: {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязателен",
		},
		KeyPriceRequired: {
			EN: "Price is required",
			RU: "Цена обязательна",
		},
		KeySalaryRequired: {
			EN: "Salary is required",
			RU: "Зарплата обязательна",
		},
		KeyContactInfoRequired: {
			EN: "Contact information is required",
			RU: "Контактная информация обязательна",
		},
		KeyPropertyTypeRequired: {
			EN: "Property type is required",
			RU: "Тип недвижимости обязателен",
		},
		KeyRoomCountRequired: {
			EN: "Room count is required",
			RU: "Количество комнат обязательно",
		},
		KeyFloorRequired: {
			EN: "Floor is required",
			RU: "Этаж обязателен",
		},
		KeyTotalFloorsRequired: {
			EN: "Total floors is required",
			RU: "Общее количество этажей обязательно",
		},
		KeyAreaRequired: {
			EN: "Area is required",
			RU: "Площадь обязательна",
		},
		KeyExperienceRequired: {
			EN: "Experience is required",
			RU: "Опыт работы обязателен",
		},
		KeyEducationRequired: {
			EN: "Education is required",
			RU: "Образование обязательно",
		},
		KeySkillsRequired: {
			EN: "Skills are required",
			RU: "Навыки обязательны",
		},
		KeyCoverLetterRequired: {
			EN: "Cover letter is required",
			RU: "Сопроводительное письмо обязательно",
		},
		KeyMessageRequired: {
			EN: "Message is required",
			RU: "Сообщение обязательно",
		},
		KeyNotificationTypeRequired: {
			EN: "Notification type is required",
			RU: "Тип уведомления обязателен",
		},
		KeyNotificationEnabledRequired: {
			EN: "Notification enabled status is required",
			RU: "Статус включения уведомлений обязателен",
		},
		KeySearchQueryRequired: {
			EN: "Search query is required",
			RU: "Поисковый запрос обязателен",
		},

		// Универсальные ID
		KeyIDRequired: {
			EN: "ID is required",
			RU: "ID обязателен",
		},
		KeyUserIDRequired: {
			EN: "User ID is required",
			RU: "ID пользователя обязателен",
		},
		KeyGroupIDRequired: {
			EN: "Group ID is required",
			RU: "ID группы обязателен",
		},

		// Универсальные операции
		KeyItemCreatedSuccess: {
			EN: "Item created successfully",
			RU: "Элемент успешно создан",
		},
		KeyItemUpdatedSuccess: {
			EN: "Item updated successfully",
			RU: "Элемент успешно обновлен",
		},
		KeyItemDeletedSuccess: {
			EN: "Item deleted successfully",
			RU: "Элемент успешно удален",
		},
		KeyItemNotFound: {
			EN: "Item not found",
			RU: "Элемент не найден",
		},
		KeyItemCreationError: {
			EN: "Item creation error",
			RU: "Ошибка создания элемента",
		},
		KeyItemUpdateError: {
			EN: "Item update error",
			RU: "Ошибка обновления элемента",
		},
		KeyItemDeleteError: {
			EN: "Item delete error",
			RU: "Ошибка удаления элемента",
		},
		KeyItemCreationFailed: {
			EN: "Failed to create item",
			RU: "Не удалось создать элемент",
		},
		KeyItemUpdateFailed: {
			EN: "Failed to update item",
			RU: "Не удалось обновить элемент",
		},
		KeyItemDeleteFailed: {
			EN: "Failed to delete item",
			RU: "Не удалось удалить элемент",
		},

		// Аутентификация и регистрация
		KeyRegistrationError: {
			EN: "Registration error",
			RU: "Ошибка регистрации",
		},
		KeyRegistrationCompletedSuccessfully: {
			EN: "Registration completed successfully",
			RU: "Регистрация выполнена успешно",
		},
		KeyUserNotFound: {
			EN: "User not found",
			RU: "Пользователь не найден",
		},
		KeyEmailNotConfirmed: {
			EN: "Email not confirmed",
			RU: "Email не подтвержден",
		},
		KeyInvalidLoginOrPassword: {
			EN: "Invalid login or password",
			RU: "Неверный логин или пароль",
		},
		KeyTokenGenerationError: {
			EN: "Token generation error",
			RU: "Ошибка генерации токена",
		},
		KeyRefreshTokenInvalid: {
			EN: "Invalid or expired refresh token",
			RU: "Недействительный или истекший refresh токен",
		},
		KeyRefreshTokenConvertError: {
			EN: "Error converting refresh token to user ID",
			RU: "Ошибка конвертации refresh токена в ID пользователя",
		},
		KeyJWTTokenCreationError: {
			EN: "Error creating JWT token",
			RU: "Ошибка создания JWT токена",
		},
		KeyUserAlreadyExists: {
			EN: "User already exists",
			RU: "Пользователь уже существует",
		},
		KeyPasswordHashError: {
			EN: "Error hashing password",
			RU: "Ошибка хеширования пароля",
		},
		KeyUserCreationError: {
			EN: "Error creating user",
			RU: "Ошибка создания пользователя",
		},
		KeyActivationCodeInvalid: {
			EN: "Invalid activation code",
			RU: "Недействительный код активации",
		},
		KeyEmailConfirmationSuccess: {
			EN: "Email confirmed successfully",
			RU: "Email успешно подтвержден",
		},
		KeyPasswordRestoreError: {
			EN: "Error restoring password",
			RU: "Ошибка восстановления пароля",
		},
		KeyPasswordRestoreSuccess: {
			EN: "Password restore email sent",
			RU: "Email для восстановления пароля отправлен",
		},
		KeyUserRoleUpdateError: {
			EN: "Error updating user role",
			RU: "Ошибка обновления роли пользователя",
		},
		KeyUserRoleUpdateSuccess: {
			EN: "User role updated successfully",
			RU: "Роль пользователя успешно обновлена",
		},
		KeyUserRolesFetchError: {
			EN: "Error fetching user roles",
			RU: "Ошибка получения ролей пользователя",
		},
		KeyUserHasRoleError: {
			EN: "Error checking user role",
			RU: "Ошибка проверки роли пользователя",
		},
		KeyAuthorizationCompletedSuccessfully: {
			EN: "Authorization completed successfully",
			RU: "Авторизация выполнена успешно",
		},
		KeyUserIdentificationError: {
			EN: "User identification error",
			RU: "Ошибка идентификации пользователя",
		},
		KeyUserRolesAndPermissions: {
			EN: "User roles and permissions",
			RU: "Роли и права пользователя",
		},
		KeyInvalidToken: {
			EN: "Invalid token",
			RU: "Недействительный токен",
		},
		KeyLoginError: {
			EN: "Login error",
			RU: "Ошибка входа",
		},
		KeyKeySuccessfullyObtained: {
			EN: "Key successfully obtained",
			RU: "Ключ успешно получен",
		},
		KeyAccountConfirmationError: {
			EN: "Account confirmation error",
			RU: "Ошибка подтверждения аккаунта",
		},
		KeyInvalidEmailAddress: {
			EN: "Invalid email address",
			RU: "Недействительный адрес email",
		},
		KeyFailedToRequestPasswordReset: {
			EN: "Failed to request password reset",
			RU: "Не удалось запросить сброс пароля",
		},
		KeyPasswordUpdated: {
			EN: "Password updated",
			RU: "Пароль обновлен",
		},

		// Общие ошибки
		KeyDatabaseError: {
			EN: "Database error",
			RU: "Ошибка базы данных",
		},
		KeyValidationError: {
			EN: "Validation error",
			RU: "Ошибка валидации",
		},
		KeyUnauthorized: {
			EN: "Unauthorized",
			RU: "Не авторизован",
		},
		KeyForbidden: {
			EN: "Forbidden",
			RU: "Доступ запрещен",
		},
		KeyNotFound: {
			EN: "Not found",
			RU: "Не найдено",
		},
		KeyInternalServerError: {
			EN: "Internal server error",
			RU: "Внутренняя ошибка сервера",
		},
		KeyServiceUnavailable: {
			EN: "Service unavailable",
			RU: "Сервис недоступен",
		},
		KeyMissingRequiredParams: {
			EN: "Missing required parameters",
			RU: "Отсутствуют обязательные параметры",
		},
		KeyDataCouldNotBeReceived: {
			EN: "Data could not be received",
			RU: "Данные не удалось получить",
		},
		KeyData: {
			EN: "Data",
			RU: "Данные",
		},
		KeyInvalidLimitParam: {
			EN: "Invalid limit parameter",
			RU: "Неверный параметр лимита",
		},

		KeyDataNotFound: {
			EN: "Data not found",
			RU: "Данные не найдены",
		},
		KeySuccess: {
			EN: "Success",
			RU: "Успешно",
		},
		KeyRecordUpdatedSuccessfully: {
			EN: "Record updated successfully",
			RU: "Запись успешно обновлена",
		},
		KeyErrorDeletingData: {
			EN: "Error deleting data",
			RU: "Ошибка удаления данных",
		},
		KeyFillUserID: {
			EN: "Please fill user ID",
			RU: "Заполните ID пользователя",
		},
		KeyNotEnoughRights: {
			EN: "Not enough rights",
			RU: "Недостаточно прав",
		},
		KeyUserUpdateError: {
			EN: "User update error",
			RU: "Ошибка обновления пользователя",
		},
		KeyErrorProcessingForm: {
			EN: "Error processing form",
			RU: "Ошибка обработки формы",
		},
		KeyDataMissingInDataField: {
			EN: "Data missing in data field",
			RU: "Отсутствуют данные в поле data",
		},
		KeyIncorrectDataInDataField: {
			EN: "Incorrect data in data field",
			RU: "Некорректные данные в поле data",
		},
		KeyInvalidInput: {
			EN: "Invalid input",
			RU: "Некорректный ввод",
		},
		KeyNoData: {
			EN: "No data",
			RU: "Нет данных",
		},
		KeyNoDataForCreation: {
			EN: "No data for creation",
			RU: "Нет данных для создания",
		},
		KeyFailedToGenerateID: {
			EN: "Failed to generate ID",
			RU: "Не удалось сгенерировать ID",
		},
		KeyUserIDShouldNotBeNil: {
			EN: "User ID should not be nil",
			RU: "ID пользователя не должен быть пустым",
		},
		KeyIDShouldNotBeNil: {
			EN: "ID should not be nil",
			RU: "ID не должен быть пустым",
		},

		KeyUserRoleNotFound: {
			EN: "User role not found",
			RU: "Роль пользователя не найдена",
		},
		KeyNoAccess: {
			EN: "No access",
			RU: "Нет доступа",
		},
		KeyIncorrectJSONData: {
			EN: "Incorrect JSON data",
			RU: "Некорректные JSON данные",
		},

		// Резюме
		KeyResumes: {
			EN: "Resumes",
			RU: "Резюме",
		},
		KeyResume: {
			EN: "Resume",
			RU: "Резюме",
		},
		KeyResumeNotFound: {
			EN: "Resume not found",
			RU: "Резюме не найдено",
		},

		KeyResumeIDShouldNotBeNil: {
			EN: "Resume ID should not be nil",
			RU: "ID резюме не должен быть пустым",
		},
		KeyResumeDoesNotBelongToUser: {
			EN: "Resume does not belong to user",
			RU: "Резюме не принадлежит пользователю",
		},
		KeyNoActiveResume: {
			EN: "No active resume",
			RU: "Нет активного резюме",
		},
		KeyResumeRequired: {
			EN: "Resume is required",
			RU: "Резюме обязательно",
		},

		// Вакансии
		KeyVacancies: {
			EN: "Vacancies",
			RU: "Вакансии",
		},
		KeyVacancy: {
			EN: "Vacancy",
			RU: "Вакансия",
		},
		KeyVacancySuccessfullyAdded: {
			EN: "Vacancy successfully added",
			RU: "Вакансия успешно добавлена",
		},
		KeyVacancyIDShouldNotBeNil: {
			EN: "Vacancy ID should not be nil",
			RU: "ID вакансии не должен быть пустым",
		},
		KeyPhone: {
			EN: "Phone",
			RU: "Телефон",
		},
		KeyNoPhone: {
			EN: "No phone",
			RU: "Нет телефона",
		},

		// Отклики на вакансии
		KeyFailedToRespond: {
			EN: "Failed to respond",
			RU: "Не удалось откликнуться",
		},
		KeyRespondedSuccessfully: {
			EN: "Responded successfully",
			RU: "Успешно откликнулись",
		},
		KeyVacancyResponds: {
			EN: "Vacancy responds",
			RU: "Отклики на вакансию",
		},
		KeyVacancyRespond: {
			EN: "Vacancy respond",
			RU: "Отклик на вакансию",
		},

		// Ошибки мессенджеров
		KeyPhoneImageGenerationError: {
			EN: "Phone image generation error",
			RU: "Ошибка генерации изображения телефона",
		},
		KeyErrorSendingToMessengers: {
			EN: "Error sending to messengers",
			RU: "Ошибка при отправке в мессенджеры",
		},
		KeyOrdersFetchError: {
			EN: "Error fetching orders",
			RU: "Ошибка получения заказов",
		},
		KeyOrderUpdateError: {
			EN: "Order update error",
			RU: "Ошибка обновления заказа",
		},
		KeySerializationError: {
			EN: "Serialization error",
			RU: "Ошибка сериализации",
		},
		KeyLocationError: {
			EN: "location or ID is nil",
			RU: "локация или ID пустые",
		},
		KeyMarshalError: {
			EN: "Error marshaling message body",
			RU: "Ошибка маршалинга тела сообщения",
		},
		KeyGroupIDInvalidFormat: {
			EN: "Group ID has invalid format",
			RU: "Group ID имеет неверный формат",
		},

		// Поиск
		KeyQueryTooShort: {
			EN: "Query is too short",
			RU: "Запрос слишком короткий",
		},
		KeySearchResult: {
			EN: "Search result",
			RU: "Результат поиска",
		},
		KeyInvalidMetaParameter: {
			EN: "Invalid meta parameter",
			RU: "Неверный параметр meta",
		},

		// Избранное

		// Мессенджеры
		KeyMessengerGetError: {
			EN: "Failed to get messengers",
			RU: "Не удалось получить мессенджеры",
		},
		KeyFilterBuildError: {
			EN: "Filter build error",
			RU: "Ошибка построения фильтра",
		},
		KeyDataProcessingError: {
			EN: "Data processing error",
			RU: "Ошибка обработки данных",
		},
		KeyDatabaseDataError: {
			EN: "Database data error",
			RU: "Ошибка получения данных из БД",
		},
		KeyNameRequired: {
			EN: "Name is required",
			RU: "Имя обязательно",
		},
		KeySkillNameRequired: {
			EN: "Skill name is required",
			RU: "Название навыка обязательно",
		},
		KeyInsertError: {
			EN: "Insert error",
			RU: "Ошибка вставки",
		},
		KeySkillDeleteIDRequired: {
			EN: "Neither ID nor binding_id provided for deletion",
			RU: "Не передан ни id, ни binding_id для удаления",
		},

		// Валюта
		KeyCurrencyNotFound: {
			EN: "Currency not found",
			RU: "Валюта не найдена",
		},

		// Пользователи
		KeyUserIDEmpty: {
			EN: "User ID cannot be empty",
			RU: "ID пользователя не может быть пустым",
		},

		// Отклики на вакансии
		KeyVacancyRespondIDRequired: {
			EN: "Vacancy respond ID is required",
			RU: "ID отклика на вакансию обязателен",
		},

		// Заказы пользователей
		KeyUserOrderIDRequired: {
			EN: "User order ID is required",
			RU: "ID заказа пользователя обязателен",
		},

		// JWT токены
		KeyInvalidSigningMethod: {
			EN: "Invalid signing method",
			RU: "Неверный метод подписи",
		},
		KeyTokenClaimsError: {
			EN: "Token claims are not of type *tokenClaims",
			RU: "Токен не является типом *tokenClaims",
		},

		// Дополнительные переводы для строковых ключей
		KeyNoFilesInForm: {
			EN: "No files in form",
			RU: "Нет файлов в форме",
		},
		KeyDirectoryCreationError: {
			EN: "Directory creation error",
			RU: "Ошибка создания директории",
		},
		KeyDiskSpaceInsufficient: {
			EN: "Insufficient disk space",
			RU: "Недостаточно места на диске",
		},
		KeyFileOpenError: {
			EN: "File open error",
			RU: "Ошибка открытия файла",
		},
		KeyFileDeletionError: {
			EN: "File deletion error",
			RU: "Ошибка удаления файла",
		},
		KeyDiskSpaceCheckError: {
			EN: "Disk space check error",
			RU: "Ошибка проверки места на диске",
		},
		KeyDiskSpaceInsufficientDetailed: {
			EN: "Insufficient disk space for operation",
			RU: "Недостаточно места на диске для операции",
		},
		KeyFilesystemStatsError: {
			EN: "Filesystem stats error",
			RU: "Ошибка статистики файловой системы",
		},
		KeyInvalidPathError: {
			EN: "Invalid path error",
			RU: "Ошибка неверного пути",
		},
		KeyDiskInfoError: {
			EN: "Disk info error",
			RU: "Ошибка информации о диске",
		},
		KeyLocationGetError: {
			EN: "Location get error",
			RU: "Ошибка получения местоположения",
		},
		KeySiteCategoryGetError: {
			EN: "Site category get error",
			RU: "Ошибка получения категории сайта",
		},
		KeyQuickAnnouncementIDShouldNotBeNil: {
			EN: "Quick announcement ID should not be nil",
			RU: "ID быстрого объявления не должен быть пустым",
		},
		KeyQuickAnnouncementRespondIDShouldNotBeNil: {
			EN: "Quick announcement respond ID should not be nil",
			RU: "ID ответа на быстрое объявление не должен быть пустым",
		},
		KeyQuickAnnouncementRequestIsNil: {
			EN: "Quick announcement request is nil",
			RU: "Запрос быстрого объявления пустой",
		},
		KeyQuickAnnouncementRespondRequestIsNil: {
			EN: "Quick announcement respond request is nil",
			RU: "Запрос ответа на быстрое объявление пустой",
		},
		KeyUUIDGenerationError: {
			EN: "UUID generation error",
			RU: "Ошибка генерации UUID",
		},
		KeyTxIsNil: {
			EN: "Transaction is nil",
			RU: "Транзакция пустая",
		},
		KeyFilterQueryIsNil: {
			EN: "Filter query is nil",
			RU: "Запрос фильтра пустой",
		},
		KeyIDParameterIsRequired: {
			EN: "ID parameter is required",
			RU: "Параметр ID обязателен",
		},
		KeyRealty: {
			EN: "Realty",
			RU: "Недвижимость",
		},
		KeyUserRoleSetFailed: {
			EN: "User role set failed",
			RU: "Не удалось установить роль пользователя",
		},
		KeyPasswordUpdateError: {
			EN: "Password update error",
			RU: "Ошибка обновления пароля",
		},
		KeyTransactionCommitError: {
			EN: "Transaction commit error",
			RU: "Ошибка фиксации транзакции",
		},
		KeyUserSearchError: {
			EN: "User search error",
			RU: "Ошибка поиска пользователя",
		},
		KeyUserRolesPermissionsError: {
			EN: "User roles permissions error",
			RU: "Ошибка прав ролей пользователя",
		},
		KeyItemTypeIDRequired: {
			EN: "Item type ID is required",
			RU: "ID типа элемента обязателен",
		},
		KeyUsernameAlreadyRegistered: {
			EN: "Username already registered",
			RU: "Имя пользователя уже зарегистрировано",
		},
		KeyCategoryNameError: {
			EN: "Category name error",
			RU: "Ошибка названия категории",
		},
		KeyUserRegistrationFailed: {
			EN: "User registration failed",
			RU: "Регистрация пользователя не удалась",
		},
		KeyActivationCodeNotFound: {
			EN: "Activation code not found",
			RU: "Код активации не найден",
		},
		KeyOfferingIDRequired: {
			EN: "Offering ID is required",
			RU: "ID предложения обязателен",
		},
		KeyEmailConfirmationError: {
			EN: "Email confirmation error",
			RU: "Ошибка подтверждения email",
		},
		KeyEmailValidationError: {
			EN: "Email validation error",
			RU: "Ошибка валидации email",
		},
		KeyTransactionCreationError: {
			EN: "Transaction creation error",
			RU: "Ошибка создания транзакции",
		},
		KeyUserFilterError: {
			EN: "User filter error",
			RU: "Ошибка фильтра пользователя",
		},
		KeyEmailAlreadyRegistered: {
			EN: "Email already registered",
			RU: "Email уже зарегистрирован",
		},
		KeyInvalidActivationCode: {
			EN: "Invalid activation code",
			RU: "Неверный код активации",
		},
		KeyOfferings: {
			EN: "Offerings",
			RU: "Предложения",
		},
		KeyLocations: {
			EN: "Locations",
			RU: "Местоположения",
		},
		KeySiteCategories: {
			EN: "Site categories",
			RU: "Категории сайта",
		},
		KeySettings: {
			EN: "Settings",
			RU: "Настройки",
		},
	}
	mu sync.RWMutex
)

// Translate возвращает перевод по ключу и языку
func Translate(key string, lang Language) string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := translations[key]; ok {
		if tr, ok := val[lang]; ok {
			return tr
		}
		if tr, ok := val[RU]; ok { // Default to Russian
			return tr
		}
	}
	return key
}
