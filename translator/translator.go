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
