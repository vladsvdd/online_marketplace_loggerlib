package translator

import (
	"sync"
)

type Language string

const (
	EN Language = "en"
	RU Language = "ru"
)

var (
	translations = map[string]map[Language]string{
		"input_data_incorrect_format": {
			EN: "Input data is in an incorrect format",
			RU: "Данные введены в неверном формате",
		},
		"username_required": {
			EN: "Username is required",
			RU: "Имя пользователя обязательно",
		},
		"username_alphanum": {
			EN: "Username must consist of letters and numbers",
			RU: "Имя пользователя должно состоять из английских букв и цифр",
		},
		"email_required": {
			EN: "Email is required",
			RU: "Email обязателен",
		},
		"email_invalid": {
			EN: "Invalid email format",
			RU: "Неверный формат email",
		},
		"password_required": {
			EN: "Password is required",
			RU: "Пароль обязателен",
		},
		"password_min_length": {
			EN: "Password must be at least 8 characters",
			RU: "Пароль должен быть не менее 8 символов",
		},
		"password_max_length": {
			EN: "Password must be no more than 32 characters",
			RU: "Пароль должен быть не более 32 символов",
		},
		"password_confirm_required": {
			EN: "Password confirmation is required",
			RU: "Подтверждение пароля обязательно",
		},
		"password_confirm_mismatch": {
			EN: "Password confirmation does not match",
			RU: "Подтверждение пароля не совпадает",
		},
		"registration_error": {
			EN: "Registration error",
			RU: "Ошибка регистрации",
		},
		"registration_completed_successfully": {
			EN: "Registration completed successfully",
			RU: "Регистрация выполнена успешно",
		},
		"user_not_found": {
			EN: "User not found",
			RU: "Пользователь не найден",
		},
		"email_not_confirmed": {
			EN: "Email not confirmed",
			RU: "Email не подтвержден",
		},
		"invalid_login_or_password": {
			EN: "Invalid login or password",
			RU: "Неверный логин или пароль",
		},
		"token_generation_error": {
			EN: "Token generation error",
			RU: "Ошибка генерации токена",
		},
		"refresh_token_invalid": {
			EN: "Invalid or expired refresh token",
			RU: "Недействительный или истекший refresh токен",
		},
		"refresh_token_convert_error": {
			EN: "Error converting refresh token to user ID",
			RU: "Ошибка конвертации refresh токена в ID пользователя",
		},
		"jwt_token_creation_error": {
			EN: "Error creating JWT token",
			RU: "Ошибка создания JWT токена",
		},
		"user_already_exists": {
			EN: "User already exists",
			RU: "Пользователь уже существует",
		},
		"password_hash_error": {
			EN: "Error hashing password",
			RU: "Ошибка хеширования пароля",
		},
		"id_generation_error": {
			EN: "Error generating user ID",
			RU: "Ошибка генерации ID пользователя",
		},
		"user_creation_error": {
			EN: "Error creating user",
			RU: "Ошибка создания пользователя",
		},
		"activation_code_invalid": {
			EN: "Invalid activation code",
			RU: "Недействительный код активации",
		},
		"email_confirmation_success": {
			EN: "Email confirmed successfully",
			RU: "Email успешно подтвержден",
		},
		"password_restore_error": {
			EN: "Error restoring password",
			RU: "Ошибка восстановления пароля",
		},
		"password_restore_success": {
			EN: "Password restore email sent",
			RU: "Email для восстановления пароля отправлен",
		},
		"user_role_update_error": {
			EN: "Error updating user role",
			RU: "Ошибка обновления роли пользователя",
		},
		"user_role_update_success": {
			EN: "User role updated successfully",
			RU: "Роль пользователя успешно обновлена",
		},
		"user_roles_fetch_error": {
			EN: "Error fetching user roles",
			RU: "Ошибка получения ролей пользователя",
		},
		"user_has_role_error": {
			EN: "Error checking user role",
			RU: "Ошибка проверки роли пользователя",
		},
		"location_required": {
			EN: "Location is required",
			RU: "Местоположение обязательно",
		},
		"site_category_required": {
			EN: "Site category is required",
			RU: "Категория сайта обязательна",
		},
		"description_required": {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		"user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязателен",
		},
		"quick_announcement_site_category_required": {
			EN: "Quick announcement site category is required",
			RU: "Категория быстрого объявления обязательна",
		},
		"title_required": {
			EN: "Title is required",
			RU: "Заголовок обязателен",
		},
		"salary_required": {
			EN: "Salary is required",
			RU: "Зарплата обязательна",
		},
		"company_id_required": {
			EN: "Company ID is required",
			RU: "ID компании обязателен",
		},
		"price_required": {
			EN: "Price is required",
			RU: "Цена обязательна",
		},
		"contact_info_required": {
			EN: "Contact information is required",
			RU: "Контактная информация обязательна",
		},
		"property_type_required": {
			EN: "Property type is required",
			RU: "Тип недвижимости обязателен",
		},
		"room_count_required": {
			EN: "Room count is required",
			RU: "Количество комнат обязательно",
		},
		"floor_required": {
			EN: "Floor is required",
			RU: "Этаж обязателен",
		},
		"total_floors_required": {
			EN: "Total floors is required",
			RU: "Общее количество этажей обязательно",
		},
		"area_required": {
			EN: "Area is required",
			RU: "Площадь обязательна",
		},
		"experience_required": {
			EN: "Experience is required",
			RU: "Опыт работы обязателен",
		},
		"education_required": {
			EN: "Education is required",
			RU: "Образование обязательно",
		},
		"skills_required": {
			EN: "Skills are required",
			RU: "Навыки обязательны",
		},
		"vacancy_id_required": {
			EN: "Vacancy ID is required",
			RU: "ID вакансии обязателен",
		},
		"cover_letter_required": {
			EN: "Cover letter is required",
			RU: "Сопроводительное письмо обязательно",
		},
		"resume_id_required": {
			EN: "Resume ID is required",
			RU: "ID резюме обязателен",
		},
		"category_id_required": {
			EN: "Category ID is required",
			RU: "ID категории обязателен",
		},
		"name_required": {
			EN: "Name is required",
			RU: "Имя обязательно",
		},
		"phone_required": {
			EN: "Phone is required",
			RU: "Телефон обязателен",
		},
		"message_required": {
			EN: "Message is required",
			RU: "Сообщение обязательно",
		},
		"chat_id_required": {
			EN: "Chat ID is required",
			RU: "ID чата обязателен",
		},
		"participant_id_required": {
			EN: "Participant ID is required",
			RU: "ID участника обязателен",
		},
		"notification_type_required": {
			EN: "Notification type is required",
			RU: "Тип уведомления обязателен",
		},
		"notification_enabled_required": {
			EN: "Notification enabled status is required",
			RU: "Статус включения уведомлений обязателен",
		},
		"search_query_required": {
			EN: "Search query is required",
			RU: "Поисковый запрос обязателен",
		},
		"favorite_added_success": {
			EN: "Added to favorites successfully",
			RU: "Успешно добавлено в избранное",
		},
		"favorite_removed_success": {
			EN: "Removed from favorites successfully",
			RU: "Успешно удалено из избранного",
		},
		"favorite_already_exists": {
			EN: "Already in favorites",
			RU: "Уже в избранном",
		},
		"favorite_not_found": {
			EN: "Not found in favorites",
			RU: "Не найдено в избранном",
		},
		"subscription_created_success": {
			EN: "Subscription created successfully",
			RU: "Подписка успешно создана",
		},
		"subscription_updated_success": {
			EN: "Subscription updated successfully",
			RU: "Подписка успешно обновлена",
		},
		"subscription_deleted_success": {
			EN: "Subscription deleted successfully",
			RU: "Подписка успешно удалена",
		},
		"subscription_not_found": {
			EN: "Subscription not found",
			RU: "Подписка не найдена",
		},
		"item_created_success": {
			EN: "Item created successfully",
			RU: "Элемент успешно создан",
		},
		"item_updated_success": {
			EN: "Item updated successfully",
			RU: "Элемент успешно обновлен",
		},
		"item_deleted_success": {
			EN: "Item deleted successfully",
			RU: "Элемент успешно удален",
		},
		"item_not_found": {
			EN: "Item not found",
			RU: "Элемент не найден",
		},
		"database_error": {
			EN: "Database error",
			RU: "Ошибка базы данных",
		},
		"validation_error": {
			EN: "Validation error",
			RU: "Ошибка валидации",
		},
		"unauthorized": {
			EN: "Unauthorized",
			RU: "Не авторизован",
		},
		"forbidden": {
			EN: "Forbidden",
			RU: "Доступ запрещен",
		},
		"not_found": {
			EN: "Not found",
			RU: "Не найдено",
		},
		"internal_server_error": {
			EN: "Internal server error",
			RU: "Внутренняя ошибка сервера",
		},
		"service_unavailable": {
			EN: "Service unavailable",
			RU: "Сервис недоступен",
		},
		"authorization_completed_successfully": {
			EN: "Authorization completed successfully",
			RU: "Авторизация выполнена успешно",
		},
		"data_could_not_be_received": {
			EN: "Data could not be received",
			RU: "Данные не удалось получить",
		},
		"data": {
			EN: "Data",
			RU: "Данные",
		},
		"invalid_limit_param": {
			EN: "Invalid limit parameter",
			RU: "Неверный параметр лимита",
		},
		"id_required": {
			EN: "ID is required",
			RU: "ID обязателен",
		},
		"data_not_found": {
			EN: "Data not found",
			RU: "Данные не найдены",
		},
		"success": {
			EN: "Success",
			RU: "Успешно",
		},
		"record_updated_successfully": {
			EN: "Record updated successfully",
			RU: "Запись успешно обновлена",
		},
		"error_deleting_data": {
			EN: "Error deleting data",
			RU: "Ошибка удаления данных",
		},
		"fill_user_id": {
			EN: "Please fill user ID",
			RU: "Заполните ID пользователя",
		},
		"not_enough_rights": {
			EN: "Not enough rights",
			RU: "Недостаточно прав",
		},
		"user_update_error": {
			EN: "User update error",
			RU: "Ошибка обновления пользователя",
		},
		"settings": {
			EN: "Settings",
			RU: "Настройки",
		},
		"error_processing_form": {
			EN: "Error processing form",
			RU: "Ошибка обработки формы",
		},
		"data_missing_in_data_field": {
			EN: "Data missing in data field",
			RU: "Отсутствуют данные в поле data",
		},
		"incorrect_data_in_data_field": {
			EN: "Incorrect data in data field",
			RU: "Некорректные данные в поле data",
		},
		"invalid_input": {
			EN: "Invalid input",
			RU: "Некорректный ввод",
		},
		"resume_does_not_belong_to_user": {
			EN: "Resume does not belong to user",
			RU: "Резюме не принадлежит пользователю",
		},
		"no_active_resume": {
			EN: "No active resume",
			RU: "Нет активного резюме",
		},
		"resume_required": {
			EN: "Resume is required",
			RU: "Резюме обязательно",
		},
		"failed_to_respond": {
			EN: "Failed to respond",
			RU: "Не удалось откликнуться",
		},
		"responded_successfully": {
			EN: "Responded successfully",
			RU: "Успешно откликнулись",
		},
		"vacancy_responds": {
			EN: "Vacancy responds",
			RU: "Отклики на вакансию",
		},
		"id_parameter_is_required": {
			EN: "ID parameter is required",
			RU: "Параметр ID обязателен",
		},
		"user_id_is_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязателен",
		},
		"vacancy_respond": {
			EN: "Vacancy respond",
			RU: "Отклик на вакансию",
		},
		"user_role_not_found": {
			EN: "User role not found",
			RU: "Роль пользователя не найдена",
		},
		"vacancies": {
			EN: "Vacancies",
			RU: "Вакансии",
		},
		"phone": {
			EN: "Phone",
			RU: "Телефон",
		},
		"vacancy_successfully_added": {
			EN: "Vacancy successfully added",
			RU: "Вакансия успешно добавлена",
		},
		"no_access": {
			EN: "No access",
			RU: "Нет доступа",
		},
		"incorrect_json_data": {
			EN: "Incorrect JSON data",
			RU: "Некорректные JSON данные",
		},
		// Ключи из auth_handler
		"username_must_consist_of_letters_and_numbers": {
			EN: "Username must consist of letters and numbers",
			RU: "Имя пользователя должно состоять из букв и цифр",
		},
		"user_identification_error": {
			EN: "User identification error",
			RU: "Ошибка идентификации пользователя",
		},
		"user_roles_and_permissions": {
			EN: "User roles and permissions",
			RU: "Роли и права пользователя",
		},
		"invalid_token": {
			EN: "Invalid token",
			RU: "Недействительный токен",
		},
		"login_error": {
			EN: "Login error",
			RU: "Ошибка входа",
		},
		"key_successfully_obtained": {
			EN: "Key successfully obtained",
			RU: "Ключ успешно получен",
		},
		"invalid_activation_code": {
			EN: "Invalid activation code",
			RU: "Недействительный код активации",
		},
		"account_confirmation_error": {
			EN: "Account confirmation error",
			RU: "Ошибка подтверждения аккаунта",
		},
		"invalid_email_address": {
			EN: "Invalid email address",
			RU: "Недействительный адрес email",
		},
		"failed_to_request_password_reset": {
			EN: "Failed to request password reset",
			RU: "Не удалось запросить сброс пароля",
		},
		"password_updated": {
			EN: "Password updated",
			RU: "Пароль обновлен",
		},
		"resumes": {
			EN: "Resumes",
			RU: "Резюме",
		},
		"resume": {
			EN: "Resume",
			RU: "Резюме",
		},
		"resume_creation_error": {
			EN: "Resume creation error",
			RU: "Ошибка создания резюме",
		},
		"resume_update_error": {
			EN: "Resume update error",
			RU: "Ошибка обновления резюме",
		},
		"resume_delete_error": {
			EN: "Resume delete error",
			RU: "Ошибка удаления резюме",
		},
		"resume_not_found": {
			EN: "Resume not found",
			RU: "Резюме не найдено",
		},
		"resume_title_required": {
			EN: "Resume title is required",
			RU: "Название резюме обязательно",
		},
		"resume_location_required": {
			EN: "Resume location is required",
			RU: "Локация резюме обязательна",
		},
		"resume_phone_required": {
			EN: "Resume phone is required",
			RU: "Телефон резюме обязателен",
		},
		"resume_description_required": {
			EN: "Resume description is required",
			RU: "Описание резюме обязательно",
		},
		"resume_created_success": {
			EN: "Resume created successfully",
			RU: "Резюме успешно создано",
		},
		"resume_updated_success": {
			EN: "Resume updated successfully",
			RU: "Резюме успешно обновлено",
		},
		"resume_deleted_success": {
			EN: "Resume deleted successfully",
			RU: "Резюме успешно удалено",
		},
		"no_data_for_creation": {
			EN: "No data for creation",
			RU: "Нет данных для создания",
		},
		"uuid_generation_error": {
			EN: "UUID generation error",
			RU: "Ошибка генерации UUID",
		},
		"currency_not_found": {
			EN: "Currency not found",
			RU: "Валюта не найдена",
		},
		"resume_creation_failed": {
			EN: "Failed to create resume",
			RU: "Не удалось создать резюме",
		},
		// Ключи для вакансий
		"vacancy": {
			EN: "Vacancy",
			RU: "Вакансия",
		},
		"vacancy_creation_error": {
			EN: "Vacancy creation error",
			RU: "Ошибка создания вакансии",
		},
		"vacancy_update_error": {
			EN: "Vacancy update error",
			RU: "Ошибка обновления вакансии",
		},
		"vacancy_delete_error": {
			EN: "Vacancy delete error",
			RU: "Ошибка удаления вакансии",
		},
		"vacancy_not_found": {
			EN: "Vacancy not found",
			RU: "Вакансия не найдена",
		},
		"vacancy_title_required": {
			EN: "Vacancy title is required",
			RU: "Название вакансии обязательно",
		},
		"vacancy_location_required": {
			EN: "Vacancy location is required",
			RU: "Локация вакансии обязательна",
		},
		"vacancy_phone_required": {
			EN: "Vacancy phone is required",
			RU: "Телефон вакансии обязателен",
		},
		"vacancy_description_required": {
			EN: "Vacancy description is required",
			RU: "Описание вакансии обязательно",
		},
		"vacancy_created_success": {
			EN: "Vacancy created successfully",
			RU: "Вакансия успешно создана",
		},
		"vacancy_updated_success": {
			EN: "Vacancy updated successfully",
			RU: "Вакансия успешно обновлена",
		},
		"vacancy_deleted_success": {
			EN: "Vacancy deleted successfully",
			RU: "Вакансия успешно удалена",
		},
		"vacancy_creation_failed": {
			EN: "Failed to create vacancy",
			RU: "Не удалось создать вакансию",
		},
		"data_processing_error": {
			EN: "Data processing error",
			RU: "Ошибка обработки данных",
		},
		"database_data_error": {
			EN: "Database data error",
			RU: "Ошибка получения данных из БД",
		},
		// Ключи для пользователей
		"user": {
			EN: "User",
			RU: "Пользователь",
		},
		"user_id_empty": {
			EN: "User ID cannot be empty",
			RU: "ID пользователя не может быть пустым",
		},
		"user_delete_error": {
			EN: "User delete error",
			RU: "Ошибка удаления пользователя",
		},
		// Ключи для недвижимости
		"realty": {
			EN: "Realty",
			RU: "Недвижимость",
		},
		"realty_creation_error": {
			EN: "Realty creation error",
			RU: "Ошибка создания недвижимости",
		},
		"realty_update_error": {
			EN: "Realty update error",
			RU: "Ошибка обновления недвижимости",
		},
		"realty_delete_error": {
			EN: "Realty delete error",
			RU: "Ошибка удаления недвижимости",
		},
		"realty_not_found": {
			EN: "Realty not found",
			RU: "Недвижимость не найдена",
		},
		"realty_user_id_required": {
			EN: "User ID is required for realty",
			RU: "ID пользователя обязателен для недвижимости",
		},
		"realty_phone_required": {
			EN: "Phone is required for realty",
			RU: "Телефон обязателен для недвижимости",
		},
		"realty_city_required": {
			EN: "City is required for realty",
			RU: "Город обязателен для недвижимости",
		},
		"realty_title_required": {
			EN: "Title is required for realty",
			RU: "Название обязательно для недвижимости",
		},
		"realty_type_required": {
			EN: "Type is required for realty",
			RU: "Тип обязателен для недвижимости",
		},
		"realty_property_type_required": {
			EN: "Property type is required for realty",
			RU: "Тип недвижимости обязателен",
		},
		"realty_description_required": {
			EN: "Description is required for realty",
			RU: "Описание обязательно для недвижимости",
		},
		"realty_created_success": {
			EN: "Realty created successfully",
			RU: "Недвижимость успешно создана",
		},
		"realty_updated_success": {
			EN: "Realty updated successfully",
			RU: "Недвижимость успешно обновлена",
		},
		"realty_deleted_success": {
			EN: "Realty deleted successfully",
			RU: "Недвижимость успешно удалена",
		},
		"realty_creation_failed": {
			EN: "Failed to create realty",
			RU: "Не удалось создать недвижимость",
		},
		// Ключи для услуг
		"offering": {
			EN: "Offering",
			RU: "Услуга",
		},
		"offering_creation_error": {
			EN: "Offering creation error",
			RU: "Ошибка создания услуги",
		},
		"offering_update_error": {
			EN: "Offering update error",
			RU: "Ошибка обновления услуги",
		},
		"offering_delete_error": {
			EN: "Offering delete error",
			RU: "Ошибка удаления услуги",
		},
		"offering_not_found": {
			EN: "Offering not found",
			RU: "Услуга не найдена",
		},
		"offering_user_id_required": {
			EN: "User ID is required for offering",
			RU: "ID пользователя обязателен для услуги",
		},
		"offering_phone_required": {
			EN: "Phone is required for offering",
			RU: "Телефон обязателен для услуги",
		},
		"offering_city_required": {
			EN: "City is required for offering",
			RU: "Город обязателен для услуги",
		},
		"offering_title_required": {
			EN: "Title is required for offering",
			RU: "Название обязательно для услуги",
		},
		"offering_category_required": {
			EN: "Category is required for offering",
			RU: "Категория обязательна для услуги",
		},
		"offering_description_required": {
			EN: "Description is required for offering",
			RU: "Описание обязательно для услуги",
		},
		"offering_created_success": {
			EN: "Offering created successfully",
			RU: "Услуга успешно создана",
		},
		"offering_updated_success": {
			EN: "Offering updated successfully",
			RU: "Услуга успешно обновлена",
		},
		"offering_deleted_success": {
			EN: "Offering deleted successfully",
			RU: "Услуга успешно удалена",
		},
		"offering_creation_failed": {
			EN: "Failed to create offering",
			RU: "Не удалось создать услугу",
		},
		"offering_update_failed": {
			EN: "Failed to update offering",
			RU: "Не удалось обновить услугу",
		},
		"offering_delete_failed": {
			EN: "Failed to delete offering",
			RU: "Не удалось удалить услугу",
		},
		"category_name_error": {
			EN: "Failed to get category name",
			RU: "Не удалось получить название категории",
		},
		// Общие ключи для ошибок
		"filter_build_error": {
			EN: "Filter build error",
			RU: "Ошибка построения фильтра",
		},
		"phone_generation_error": {
			EN: "Phone image generation error",
			RU: "Ошибка генерации изображения телефона",
		},
		"user_id_should_not_be_nil": {
			EN: "User ID should not be nil",
			RU: "ID пользователя не должен быть пустым",
		},
		"id_should_not_be_nil": {
			EN: "ID should not be nil",
			RU: "ID не должен быть пустым",
		},
		"resume_id_should_not_be_nil": {
			EN: "Resume ID should not be nil",
			RU: "ID резюме не должен быть пустым",
		},
		"vacancy_id_should_not_be_nil": {
			EN: "Vacancy ID should not be nil",
			RU: "ID вакансии не должен быть пустым",
		},
		"realty_id_should_not_be_nil": {
			EN: "Realty ID should not be nil",
			RU: "ID недвижимости не должен быть пустым",
		},
		"offering_id_should_not_be_nil": {
			EN: "Offering ID should not be nil",
			RU: "ID услуги не должен быть пустым",
		},
		"quick_announcement_id_should_not_be_nil": {
			EN: "Quick announcement ID should not be nil",
			RU: "ID быстрого объявления не должен быть пустым",
		},
		"orders_error": {
			EN: "Error getting orders",
			RU: "Ошибка получения заказов",
		},
		"order_update_error": {
			EN: "Order update error",
			RU: "Ошибка обновления заказа",
		},
		"delete_error": {
			EN: "Delete error",
			RU: "Ошибка удаления",
		},
		"telegram_error": {
			EN: "Telegram error",
			RU: "Ошибка Telegram",
		},
		"vk_error": {
			EN: "VK error",
			RU: "Ошибка VK",
		},
		"group_id_invalid_format": {
			EN: "Group ID has invalid format",
			RU: "Group ID имеет неверный формат",
		},
		"telegram_marshaling_error": {
			EN: "Error marshaling message body for Telegram",
			RU: "Ошибка маршалинга тела сообщения для Telegram",
		},
		"vk_marshaling_error": {
			EN: "Error marshaling message body for VK",
			RU: "Ошибка маршалинга тела сообщения для VK",
		},
		"user_id_not_filled": {
			EN: "User ID is not filled",
			RU: "User ID не заполнен",
		},
		// Ключи для файлов
		"tx_is_nil": {
			EN: "Transaction is nil",
			RU: "Транзакция не инициализирована",
		},
		"filter_query_is_nil": {
			EN: "Filter query is nil",
			RU: "Пустой фильтр запроса",
		},
		"db_delete_error": {
			EN: "Database delete error",
			RU: "Ошибка удаления данных из БД",
		},
		"file_not_found": {
			EN: "File not found",
			RU: "Файл не найден",
		},
		"file_creation_error": {
			EN: "File creation error",
			RU: "Ошибка создания файла",
		},
		"file_update_error": {
			EN: "File update error",
			RU: "Ошибка обновления файла",
		},
		"file_delete_error": {
			EN: "File delete error",
			RU: "Ошибка удаления файла",
		},
		"file_id_required": {
			EN: "File ID is required",
			RU: "ID файла обязателен",
		},
		"file_name_required": {
			EN: "File name is required",
			RU: "Имя файла обязательно",
		},
		"file_path_required": {
			EN: "File path is required",
			RU: "Путь к файлу обязателен",
		},
		"file_type_required": {
			EN: "File type is required",
			RU: "Тип файла обязателен",
		},
		"file_size_required": {
			EN: "File size is required",
			RU: "Размер файла обязателен",
		},
		"file_upload_error": {
			EN: "File upload error",
			RU: "Ошибка загрузки файла",
		},
		"file_download_error": {
			EN: "File download error",
			RU: "Ошибка скачивания файла",
		},
		"file_permission_denied": {
			EN: "File permission denied",
			RU: "Нет прав на файл",
		},
		"file_already_exists": {
			EN: "File already exists",
			RU: "Файл уже существует",
		},
		"file_storage_error": {
			EN: "File storage error",
			RU: "Ошибка хранилища файлов",
		},
		"file_read_error": {
			EN: "File read error",
			RU: "Ошибка чтения файла",
		},
		"file_write_error": {
			EN: "File write error",
			RU: "Ошибка записи файла",
		},
		"file_remove_error": {
			EN: "File remove error",
			RU: "Ошибка удаления файла",
		},
		"file_list_error": {
			EN: "File list error",
			RU: "Ошибка получения списка файлов",
		},
		"file_meta_error": {
			EN: "File meta error",
			RU: "Ошибка метаданных файла",
		},
		"file_duplicate_error": {
			EN: "File duplicate error",
			RU: "Дубликат файла",
		},
		"file_invalid_error": {
			EN: "File invalid error",
			RU: "Некорректный файл",
		},
		"file_access_error": {
			EN: "File access error",
			RU: "Ошибка доступа к файлу",
		},
		"file_limit_exceeded": {
			EN: "File limit exceeded",
			RU: "Превышен лимит файлов",
		},
		"file_unsupported_type": {
			EN: "File unsupported type",
			RU: "Неподдерживаемый тип файла",
		},
		"file_conversion_error": {
			EN: "File conversion error",
			RU: "Ошибка конвертации файла",
		},
		"file_move_error": {
			EN: "File move error",
			RU: "Ошибка перемещения файла",
		},
		"file_copy_error": {
			EN: "File copy error",
			RU: "Ошибка копирования файла",
		},
		"file_rename_error": {
			EN: "File rename error",
			RU: "Ошибка переименования файла",
		},
		"file_lock_error": {
			EN: "File lock error",
			RU: "Ошибка блокировки файла",
		},
		"file_unlock_error": {
			EN: "File unlock error",
			RU: "Ошибка разблокировки файла",
		},
		"file_share_error": {
			EN: "File share error",
			RU: "Ошибка публикации файла",
		},
		"file_unshare_error": {
			EN: "File unshare error",
			RU: "Ошибка снятия публикации файла",
		},
		"file_tag_error": {
			EN: "File tag error",
			RU: "Ошибка добавления тега к файлу",
		},
		"file_untag_error": {
			EN: "File untag error",
			RU: "Ошибка удаления тега у файла",
		},
		"file_version_error": {
			EN: "File version error",
			RU: "Ошибка версии файла",
		},
		"file_restore_error": {
			EN: "File restore error",
			RU: "Ошибка восстановления файла",
		},
		"file_archive_error": {
			EN: "File archive error",
			RU: "Ошибка архивирования файла",
		},
		"file_unarchive_error": {
			EN: "File unarchive error",
			RU: "Ошибка разархивирования файла",
		},
		"file_sync_error": {
			EN: "File sync error",
			RU: "Ошибка синхронизации файла",
		},
		"file_unsync_error": {
			EN: "File unsync error",
			RU: "Ошибка отмены синхронизации файла",
		},
		"file_quota_exceeded": {
			EN: "File quota exceeded",
			RU: "Превышена квота файлов",
		},
		"file_expired_error": {
			EN: "File expired error",
			RU: "Срок действия файла истек",
		},
		"file_link_error": {
			EN: "File link error",
			RU: "Ошибка ссылки на файл",
		},
		"file_unlink_error": {
			EN: "File unlink error",
			RU: "Ошибка удаления ссылки на файл",
		},
		"file_checksum_error": {
			EN: "File checksum error",
			RU: "Ошибка контрольной суммы файла",
		},
		"file_corrupt_error": {
			EN: "File corrupt error",
			RU: "Файл поврежден",
		},
		"file_scan_error": {
			EN: "File scan error",
			RU: "Ошибка сканирования файла",
		},
		"file_virus_error": {
			EN: "File virus error",
			RU: "Обнаружен вирус в файле",
		},
		"file_policy_error": {
			EN: "File policy error",
			RU: "Ошибка политики файла",
		},
		"file_external_error": {
			EN: "File external error",
			RU: "Внешняя ошибка файла",
		},
		"file_internal_error": {
			EN: "File internal error",
			RU: "Внутренняя ошибка файла",
		},
		"file_unknown_error": {
			EN: "Unknown file error",
			RU: "Неизвестная ошибка файла",
		},
		// Ключи для подписок
		"subscription_request_is_nil": {
			EN: "Subscription request is nil",
			RU: "Запрос подписки не инициализирован",
		},
		"subscriber_id_required": {
			EN: "Subscriber ID is required",
			RU: "ID подписчика обязателен",
		},
		"target_user_id_required": {
			EN: "Target user ID is required",
			RU: "ID целевого пользователя обязателен",
		},
		"subscription_id_should_not_be_empty": {
			EN: "Subscription ID should not be empty",
			RU: "ID подписки не должен быть пустым",
		},
		"subscription_creation_error": {
			EN: "Subscription creation error",
			RU: "Ошибка создания подписки",
		},
		"subscription_update_error": {
			EN: "Subscription update error",
			RU: "Ошибка обновления подписки",
		},
		"subscription_delete_error": {
			EN: "Subscription delete error",
			RU: "Ошибка удаления подписки",
		},
		"subscription_creation_failed": {
			EN: "Failed to create subscription",
			RU: "Не удалось создать подписку",
		},
		"subscription_update_failed": {
			EN: "Failed to update subscription",
			RU: "Не удалось обновить подписку",
		},
		"subscription_delete_failed": {
			EN: "Failed to delete subscription",
			RU: "Не удалось удалить подписку",
		},
		"subscription_schedule_error": {
			EN: "Subscription schedule error",
			RU: "Ошибка расписания подписки",
		},
		"subscription_schedule_add_error": {
			EN: "Subscription schedule add error",
			RU: "Ошибка добавления расписания подписки",
		},
		"subscription_schedule_update_error": {
			EN: "Subscription schedule update error",
			RU: "Ошибка обновления расписания подписки",
		},
		"subscription_schedule_delete_error": {
			EN: "Subscription schedule delete error",
			RU: "Ошибка удаления расписания подписки",
		},
		"subscription_schedule_not_found": {
			EN: "Subscription schedule not found",
			RU: "Расписание подписки не найдено",
		},
		"subscription_schedule_created_success": {
			EN: "Subscription schedule created successfully",
			RU: "Расписание подписки успешно создано",
		},
		"subscription_schedule_updated_success": {
			EN: "Subscription schedule updated successfully",
			RU: "Расписание подписки успешно обновлено",
		},
		"subscription_schedule_deleted_success": {
			EN: "Subscription schedule deleted successfully",
			RU: "Расписание подписки успешно удалено",
		},
		"subscription_schedule_creation_failed": {
			EN: "Failed to create subscription schedule",
			RU: "Не удалось создать расписание подписки",
		},
		"subscription_schedule_update_failed": {
			EN: "Failed to update subscription schedule",
			RU: "Не удалось обновить расписание подписки",
		},
		"subscription_schedule_delete_failed": {
			EN: "Failed to delete subscription schedule",
			RU: "Не удалось удалить расписание подписки",
		},
		// Ключи для откликов на вакансии
		"vacancy_respond_id_required": {
			EN: "Vacancy respond ID is required",
			RU: "ID отклика на вакансию обязателен",
		},
		"vacancy_respond_creation_error": {
			EN: "Vacancy respond creation error",
			RU: "Ошибка создания отклика на вакансию",
		},
		"vacancy_respond_update_error": {
			EN: "Vacancy respond update error",
			RU: "Ошибка обновления отклика на вакансию",
		},
		"vacancy_respond_delete_error": {
			EN: "Vacancy respond delete error",
			RU: "Ошибка удаления отклика на вакансию",
		},
		"vacancy_respond_not_found": {
			EN: "Vacancy respond not found",
			RU: "Отклик на вакансию не найден",
		},
		"vacancy_respond_created_success": {
			EN: "Vacancy respond created successfully",
			RU: "Отклик на вакансию успешно создан",
		},
		"vacancy_respond_updated_success": {
			EN: "Vacancy respond updated successfully",
			RU: "Отклик на вакансию успешно обновлен",
		},
		"vacancy_respond_deleted_success": {
			EN: "Vacancy respond deleted successfully",
			RU: "Отклик на вакансию успешно удален",
		},
		"vacancy_respond_creation_failed": {
			EN: "Failed to create vacancy respond",
			RU: "Не удалось создать отклик на вакансию",
		},
		"vacancy_respond_update_failed": {
			EN: "Failed to update vacancy respond",
			RU: "Не удалось обновить отклик на вакансию",
		},
		"vacancy_respond_delete_failed": {
			EN: "Failed to delete vacancy respond",
			RU: "Не удалось удалить отклик на вакансию",
		},
		// Ключи для заказов пользователей
		"user_order_id_required": {
			EN: "User order ID is required",
			RU: "ID заказа пользователя обязателен",
		},
		"user_order_creation_error": {
			EN: "User order creation error",
			RU: "Ошибка создания заказа пользователя",
		},
		"user_order_update_error": {
			EN: "User order update error",
			RU: "Ошибка обновления заказа пользователя",
		},
		"user_order_delete_error": {
			EN: "User order delete error",
			RU: "Ошибка удаления заказа пользователя",
		},
		"user_order_not_found": {
			EN: "User order not found",
			RU: "Заказ пользователя не найден",
		},
		"user_order_created_success": {
			EN: "User order created successfully",
			RU: "Заказ пользователя успешно создан",
		},
		"user_order_updated_success": {
			EN: "User order updated successfully",
			RU: "Заказ пользователя успешно обновлен",
		},
		"user_order_deleted_success": {
			EN: "User order deleted successfully",
			RU: "Заказ пользователя успешно удален",
		},
		"user_order_creation_failed": {
			EN: "Failed to create user order",
			RU: "Не удалось создать заказ пользователя",
		},
		"user_order_update_failed": {
			EN: "Failed to update user order",
			RU: "Не удалось обновить заказ пользователя",
		},
		"user_order_delete_failed": {
			EN: "Failed to delete user order",
			RU: "Не удалось удалить заказ пользователя",
		},
		// Ключи для уведомлений пользователей
		"user_notification_id_required": {
			EN: "User notification ID is required",
			RU: "ID уведомления пользователя обязателен",
		},
		"user_notification_creation_error": {
			EN: "User notification creation error",
			RU: "Ошибка создания уведомления пользователя",
		},
		"user_notification_update_error": {
			EN: "User notification update error",
			RU: "Ошибка обновления уведомления пользователя",
		},
		"user_notification_delete_error": {
			EN: "User notification delete error",
			RU: "Ошибка удаления уведомления пользователя",
		},
		"user_notification_not_found": {
			EN: "User notification not found",
			RU: "Уведомление пользователя не найдено",
		},
		"user_notification_created_success": {
			EN: "User notification created successfully",
			RU: "Уведомление пользователя успешно создано",
		},
		"user_notification_updated_success": {
			EN: "User notification updated successfully",
			RU: "Уведомление пользователя успешно обновлено",
		},
		"user_notification_deleted_success": {
			EN: "User notification deleted successfully",
			RU: "Уведомление пользователя успешно удалено",
		},
		"user_notification_creation_failed": {
			EN: "Failed to create user notification",
			RU: "Не удалось создать уведомление пользователя",
		},
		"user_notification_update_failed": {
			EN: "Failed to update user notification",
			RU: "Не удалось обновить уведомление пользователя",
		},
		"user_notification_delete_failed": {
			EN: "Failed to delete user notification",
			RU: "Не удалось удалить уведомление пользователя",
		},
		"user_notification_request_is_nil": {
			EN: "User notification request is nil",
			RU: "Запрос уведомления пользователя не инициализирован",
		},
		// Ключи для услуг (offering)
		"offering_id_required": {
			EN: "Offering ID is required",
			RU: "ID услуги обязателен",
		},
		"offering_price_required": {
			EN: "Offering price is required",
			RU: "Цена услуги обязательна",
		},
		"category_required": {
			EN: "Category is required",
			RU: "Категория обязательна",
		},
		// Ключи для быстрых объявлений
		"quick_announcement_id_required": {
			EN: "Quick announcement ID is required",
			RU: "ID быстрого объявления обязателен",
		},
		"quick_announcement_title_required": {
			EN: "Quick announcement title is required",
			RU: "Название быстрого объявления обязательно",
		},
		"quick_announcement_description_required": {
			EN: "Quick announcement description is required",
			RU: "Описание быстрого объявления обязательно",
		},
		"quick_announcement_creation_error": {
			EN: "Quick announcement creation error",
			RU: "Ошибка создания быстрого объявления",
		},
		"quick_announcement_update_error": {
			EN: "Quick announcement update error",
			RU: "Ошибка обновления быстрого объявления",
		},
		"quick_announcement_delete_error": {
			EN: "Quick announcement delete error",
			RU: "Ошибка удаления быстрого объявления",
		},
		"quick_announcement_not_found": {
			EN: "Quick announcement not found",
			RU: "Быстрое объявление не найдено",
		},
		"quick_announcement_created_success": {
			EN: "Quick announcement created successfully",
			RU: "Быстрое объявление успешно создано",
		},
		"quick_announcement_updated_success": {
			EN: "Quick announcement updated successfully",
			RU: "Быстрое объявление успешно обновлено",
		},
		"quick_announcement_deleted_success": {
			EN: "Quick announcement deleted successfully",
			RU: "Быстрое объявление успешно удалено",
		},
		"quick_announcement_creation_failed": {
			EN: "Failed to create quick announcement",
			RU: "Не удалось создать быстрое объявление",
		},
		"quick_announcement_update_failed": {
			EN: "Failed to update quick announcement",
			RU: "Не удалось обновить быстрое объявление",
		},
		"quick_announcement_delete_failed": {
			EN: "Failed to delete quick announcement",
			RU: "Не удалось удалить быстрое объявление",
		},
		"quick_announcement_request_is_nil": {
			EN: "Quick announcement request is nil",
			RU: "Запрос быстрого объявления не инициализирован",
		},
		// Ключи для откликов на быстрые объявления
		"quick_announcement_respond_id_required": {
			EN: "Quick announcement respond ID is required",
			RU: "ID отклика на быстрое объявление обязателен",
		},
		"quick_announcement_respond_creation_error": {
			EN: "Quick announcement respond creation error",
			RU: "Ошибка создания отклика на быстрое объявление",
		},
		"quick_announcement_respond_update_error": {
			EN: "Quick announcement respond update error",
			RU: "Ошибка обновления отклика на быстрое объявление",
		},
		"quick_announcement_respond_delete_error": {
			EN: "Quick announcement respond delete error",
			RU: "Ошибка удаления отклика на быстрое объявление",
		},
		"quick_announcement_respond_not_found": {
			EN: "Quick announcement respond not found",
			RU: "Отклик на быстрое объявление не найден",
		},
		"quick_announcement_respond_created_success": {
			EN: "Quick announcement respond created successfully",
			RU: "Отклик на быстрое объявление успешно создан",
		},
		"quick_announcement_respond_updated_success": {
			EN: "Quick announcement respond updated successfully",
			RU: "Отклик на быстрое объявление успешно обновлен",
		},
		"quick_announcement_respond_deleted_success": {
			EN: "Quick announcement respond deleted successfully",
			RU: "Отклик на быстрое объявление успешно удален",
		},
		"quick_announcement_respond_creation_failed": {
			EN: "Failed to create quick announcement respond",
			RU: "Не удалось создать отклик на быстрое объявление",
		},
		"quick_announcement_respond_update_failed": {
			EN: "Failed to update quick announcement respond",
			RU: "Не удалось обновить отклик на быстрое объявление",
		},
		"quick_announcement_respond_delete_failed": {
			EN: "Failed to delete quick announcement respond",
			RU: "Не удалось удалить отклик на быстрое объявление",
		},
		"quick_announcement_respond_request_is_nil": {
			EN: "Quick announcement respond request is nil",
			RU: "Запрос отклика на быстрое объявление не инициализирован",
		},
		// Ключи для избранного
		"favourites_id_required": {
			EN: "Favourites ID is required",
			RU: "ID избранного обязателен",
		},
		"favourites_creation_error": {
			EN: "Favourites creation error",
			RU: "Ошибка создания избранного",
		},
		"favourites_update_error": {
			EN: "Favourites update error",
			RU: "Ошибка обновления избранного",
		},
		"favourites_delete_error": {
			EN: "Favourites delete error",
			RU: "Ошибка удаления избранного",
		},
		"favourites_not_found": {
			EN: "Favourites not found",
			RU: "Избранное не найдено",
		},
		"favourites_created_success": {
			EN: "Favourites created successfully",
			RU: "Избранное успешно создано",
		},
		"favourites_updated_success": {
			EN: "Favourites updated successfully",
			RU: "Избранное успешно обновлено",
		},
		"favourites_deleted_success": {
			EN: "Favourites deleted successfully",
			RU: "Избранное успешно удалено",
		},
		"favourites_creation_failed": {
			EN: "Failed to create favourites",
			RU: "Не удалось создать избранное",
		},
		"favourites_update_failed": {
			EN: "Failed to update favourites",
			RU: "Не удалось обновить избранное",
		},
		"favourites_delete_failed": {
			EN: "Failed to delete favourites",
			RU: "Не удалось удалить избранное",
		},
		"item_id_required": {
			EN: "Item ID is required",
			RU: "ID элемента обязателен",
		},
		"item_type_id_required": {
			EN: "Item type ID is required",
			RU: "ID типа элемента обязателен",
		},
		// Ключи для мессенджеров
		"messenger_id_required": {
			EN: "Messenger ID is required",
			RU: "ID мессенджера обязателен",
		},
		"messenger_creation_error": {
			EN: "Messenger creation error",
			RU: "Ошибка создания мессенджера",
		},
		"messenger_update_error": {
			EN: "Messenger update error",
			RU: "Ошибка обновления мессенджера",
		},
		"messenger_delete_error": {
			EN: "Messenger delete error",
			RU: "Ошибка удаления мессенджера",
		},
		"messenger_not_found": {
			EN: "Messenger not found",
			RU: "Мессенджер не найден",
		},
		"messenger_created_success": {
			EN: "Messenger created successfully",
			RU: "Мессенджер успешно создан",
		},
		"messenger_updated_success": {
			EN: "Messenger updated successfully",
			RU: "Мессенджер успешно обновлен",
		},
		"messenger_deleted_success": {
			EN: "Messenger deleted successfully",
			RU: "Мессенджер успешно удален",
		},
		"messenger_creation_failed": {
			EN: "Failed to create messenger",
			RU: "Не удалось создать мессенджер",
		},
		"messenger_update_failed": {
			EN: "Failed to update messenger",
			RU: "Не удалось обновить мессенджер",
		},
		"messenger_delete_failed": {
			EN: "Failed to delete messenger",
			RU: "Не удалось удалить мессенджер",
		},
		"messenger_get_error": {
			EN: "Failed to get messengers",
			RU: "Не удалось получить мессенджеры",
		},
		"group_id_required": {
			EN: "Group ID is required",
			RU: "ID группы обязателен",
		},
		"site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязателен",
		},
		"auth_id_required": {
			EN: "Auth ID is required",
			RU: "ID аутентификации обязателен",
		},
		"auth_creation_error": {
			EN: "Auth creation error",
			RU: "Ошибка создания аутентификации",
		},
		"auth_update_error": {
			EN: "Auth update error",
			RU: "Ошибка обновления аутентификации",
		},
		"auth_delete_error": {
			EN: "Auth delete error",
			RU: "Ошибка удаления аутентификации",
		},
		"auth_not_found": {
			EN: "Auth not found",
			RU: "Аутентификация не найдена",
		},
		"auth_created_success": {
			EN: "Auth created successfully",
			RU: "Аутентификация успешно создана",
		},
		"auth_updated_success": {
			EN: "Auth updated successfully",
			RU: "Аутентификация успешно обновлена",
		},
		"auth_deleted_success": {
			EN: "Auth deleted successfully",
			RU: "Аутентификация успешно удалена",
		},
		"auth_creation_failed": {
			EN: "Failed to create auth",
			RU: "Не удалось создать аутентификацию",
		},
		"auth_update_failed": {
			EN: "Failed to update auth",
			RU: "Не удалось обновить аутентификацию",
		},
		"auth_delete_failed": {
			EN: "Failed to delete auth",
			RU: "Не удалось удалить аутентификацию",
		},
		"auth_invalid_credentials": {
			EN: "Invalid credentials",
			RU: "Неверные учетные данные",
		},
		"auth_user_not_found": {
			EN: "User not found",
			RU: "Пользователь не найден",
		},
		"auth_password_mismatch": {
			EN: "Password mismatch",
			RU: "Неверный пароль",
		},
		"user_filter_error": {
			EN: "User filter error",
			RU: "Ошибка фильтрации пользователей",
		},
		"email_already_registered": {
			EN: "User with this email is already registered",
			RU: "Пользователь с таким email уже зарегистрирован",
		},
		"username_already_registered": {
			EN: "User with this username is already registered",
			RU: "Пользователь с таким username уже зарегистрирован",
		},
		"user_registration_failed": {
			EN: "Failed to register user",
			RU: "Не удалось зарегистрировать пользователя",
		},
		"user_role_set_failed": {
			EN: "Failed to set user role",
			RU: "Не удалось установить роль пользователю",
		},
		"activation_code_not_found": {
			EN: "Activation code not found",
			RU: "Код активации не найден",
		},
		"email_confirmation_error": {
			EN: "Email confirmation error",
			RU: "Ошибка подтверждения регистрации",
		},
		"user_search_error": {
			EN: "User search error",
			RU: "Ошибка поиска пользователя по email",
		},
		"email_validation_error": {
			EN: "Please check if you filled the email field correctly",
			RU: "Проверьте правильно ли вы заполнили поле email",
		},
		"transaction_creation_error": {
			EN: "Transaction creation error",
			RU: "Ошибка при создании транзакции",
		},
		"password_update_error": {
			EN: "Password update error",
			RU: "Возникла ошибка обновления пароля пользователя",
		},
		"transaction_commit_error": {
			EN: "Transaction commit error",
			RU: "Ошибка при коммите транзакции",
		},
		"user_roles_permissions_error": {
			EN: "User roles and permissions error",
			RU: "Ошибка получения ролей и разрешений пользователя",
		},
		"user_role_update_failed": {
			EN: "Failed to update user role",
			RU: "Не удалось обновить роль пользователя",
		},
		"skill_id_required": {
			EN: "Skill ID is required",
			RU: "ID навыка обязателен",
		},
		"skill_name_required": {
			EN: "Skill name is required",
			RU: "Название навыка обязательно",
		},
		"skill_insert_error": {
			EN: "Skill insert error",
			RU: "Ошибка вставки умений",
		},
		"skill_delete_id_required": {
			EN: "Neither ID nor binding_id provided for deletion",
			RU: "Не передан ни id, ни binding_id для удаления",
		},
		"skill_delete_error": {
			EN: "Skill deletion error",
			RU: "Ошибка удаления навыков",
		},
		"listing_amenities_id_required": {
			EN: "Listing amenities ID is required",
			RU: "ID удобств обязателен",
		},
		"listing_amenities_insert_error": {
			EN: "Listing amenities insert error",
			RU: "Ошибка добавления графика работы",
		},
		"listing_amenities_delete_id_required": {
			EN: "Neither ID nor binding_id provided for deletion",
			RU: "Не передан ни id, или ни binding_id для удаления",
		},
		"listing_amenities_delete_error": {
			EN: "Listing amenities deletion error",
			RU: "Ошибка удаления графика работы",
		},
		"site_category_get_error": {
			EN: "Failed to get site categories",
			RU: "Не удалось получить локации",
		},
		"location_get_error": {
			EN: "Failed to get locations",
			RU: "Не удалось получить локации",
		},
		"location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязателен",
		},
		"offering_schedule_vacancy_id_required": {
			EN: "Vacancy ID is required",
			RU: "Заполните id вакансии",
		},
		"offering_schedule_insert_error": {
			EN: "Offering schedule insert error",
			RU: "Ошибка добавления графика работы",
		},
		"offering_schedule_delete_id_required": {
			EN: "Neither ID nor offering_id provided for deletion",
			RU: "Не передан ни id, или ни offering_id для удаления",
		},
		"offering_schedule_delete_error": {
			EN: "Offering schedule deletion error",
			RU: "Ошибка удаления графика работы",
		},
		"no_data": {
			EN: "No data",
			RU: "Нет данных",
		},
		"phone_image_generation_error": {
			EN: "Phone image generation error",
			RU: "Ошибка генерации изображения телефона",
		},
		"orders_fetch_error": {
			EN: "Error fetching orders",
			RU: "Ошибка получения заказов",
		},
		"telegram_location_error": {
			EN: "Telegram: location or ID is nil",
			RU: "Telegram: локация или ID пустые",
		},
		"vk_location_error": {
			EN: "VK: location or ID is nil",
			RU: "VK: локация или ID пустые",
		},
		"telegram_marshal_error": {
			EN: "Error marshaling message body for Telegram",
			RU: "Ошибка маршалинга тела сообщения для телеграма",
		},
		"vk_marshal_error": {
			EN: "Error marshaling message body for VK",
			RU: "Ошибка маршалинга тела сообщения для vk",
		},
		"vacancy_respond_send_error": {
			EN: "Error sending vacancy respond",
			RU: "Ошибка отправки отклика, повторите позже",
		},
		"user_setting_update_error": {
			EN: "User setting update error",
			RU: "Ошибка обновления настроек пользователя",
		},
		"notification_creation_error": {
			EN: "Notification creation error",
			RU: "Ошибка создания уведомления, повторите позже",
		},
		"notification_update_error": {
			EN: "Notification update error",
			RU: "Ошибка обновления уведомления, повторите позже",
		},
		"notification_delete_error": {
			EN: "Notification deletion error",
			RU: "Ошибка удаления уведомления, повторите позже",
		},
		"site_category_fetch_error": {
			EN: "Site category fetch error",
			RU: "Ошибка получения категорий сайта, повторите позже",
		},
		"quick_announcement_respond_id_should_not_be_nil": {
			EN: "Quick announcement respond ID should not be nil",
			RU: "ID быстрого отклика не должен быть пустым",
		},
		"messenger_marshal_error": {
			EN: "Error marshaling message body for Telegram",
			RU: "Ошибка маршалинга тела сообщения для телеграма",
		},
		"failed_to_generate_id": {
			EN: "Failed to generate ID",
			RU: "Не удалось сгенерировать ID",
		},
		"error_sending_to_messengers": {
			EN: "Error sending to messengers",
			RU: "Ошибка при отправке в мессенджеры",
		},
		"serialization_error": {
			EN: "Serialization error",
			RU: "Ошибка сериализации",
		},
		"invalid_signing_method": {
			EN: "Invalid signing method",
			RU: "Неверный метод подписи",
		},
		"token_claims_error": {
			EN: "Token claims are not of type *tokenClaims",
			RU: "Токен не является типом *tokenClaims",
		},
		"no_files_in_form": {
			EN: "No files in the form",
			RU: "Файлы в форме отсутствуют",
		},
		"disk_space_insufficient": {
			EN: "Insufficient disk space",
			RU: "Недостаточно места на диске",
		},
		"binding_id_required": {
			EN: "Binding ID or ID should not be empty",
			RU: "Binding ID или ID не должны быть пустыми",
		},
		"file_deletion_error": {
			EN: "Failed to delete file",
			RU: "Не удалось удалить файл",
		},
		"disk_space_check_error": {
			EN: "Failed to check disk space",
			RU: "Не удалось проверить дисковое пространство",
		},
		"disk_space_insufficient_detailed": {
			EN: "Insufficient disk space: available %d bytes, required %d bytes",
			RU: "Недостаточно места на диске: доступно %d байт, требуется %d байт",
		},
		"invalid_path_error": {
			EN: "Invalid path",
			RU: "Неверный путь",
		},
		"disk_info_error": {
			EN: "Error getting disk information",
			RU: "Ошибка при получении информации о диске",
		},
		"filesystem_stats_error": {
			EN: "Failed to get filesystem statistics",
			RU: "Не удалось получить статистику файловой системы",
		},
		"vacancy_update_failed": {
			EN: "Failed to update vacancy",
			RU: "Не удалось обновить вакансию",
		},
		"notification_creation_failed": {
			EN: "Failed to create notification",
			RU: "Не удалось создать уведомление",
		},
		"notification_update_failed": {
			EN: "Failed to update notification",
			RU: "Не удалось обновить уведомление",
		},
		"resume_update_failed": {
			EN: "Failed to update resume",
			RU: "Не удалось обновить резюме",
		},
		"no_phone": {
			EN: "No phone",
			RU: "Нет телефона",
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
