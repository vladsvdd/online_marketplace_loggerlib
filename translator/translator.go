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
	translations = make(map[string]map[Language]string)
	mu           sync.RWMutex
)

// LoadTranslations загружает переводы из JSON-файла
func LoadTranslations() {
	mu.Lock()
	defer mu.Unlock()
	translations = map[string]map[Language]string{
		"input_data_incorrect_format": {
			EN: "Input data is in an incorrect format",
			RU: "Данные введены в неверном формате",
		},
		"username_must_consist_of_letters_and_numbers": {
			EN: "Username must consist of letters and numbers",
			RU: "Имя пользователя должно состоять из букв и цифр",
		},
		"registration_completed_successfully": {
			EN: "Registration completed successfully",
			RU: "Регистрация выполнена успешно",
		},
		"key_successfully_obtained": {
			EN: "Key received successfully",
			RU: "Ключ успешно получен",
		},
		"invalid_activation_code": {
			EN: "Activation code is invalid",
			RU: "Код активации недействителен",
		},
		"account_confirmation_error": {
			EN: "Account confirmation error",
			RU: "Ошибка подтверждения учетной записи",
		},
		"invalid_email_address": {
			EN: "Invalid email address",
			RU: "Некорректный адрес электронной почты",
		},
		"failed_to_request_password_reset": {
			EN: "Failed to request password reset",
			RU: "Не удалось запросить восстановление пароля",
		},
		"item_id_required": {
			EN: "Item ID is required",
			RU: "item_id обязателен",
		},
		"error_processing_form": {
			EN: "Error processing form",
			RU: "Ошибка при обработке формы",
		},
		"data_missing_in_data_field": {
			EN: "Data missing in 'data' field",
			RU: "Отсутствуют данные в поле 'data'",
		},
		"incorrect_data_in_data_field": {
			EN: "Incorrect data in 'data' field",
			RU: "Некорректные данные в поле 'data'",
		},
		"settings": {
			EN: "Settings",
			RU: "Настройки",
		},
		"vacancy_id_required": {
			EN: "Vacancy ID is required",
			RU: "vacancyID обязателен",
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
			RU: "Отклик успешно отправлен",
		},
		"vacancy_responds": {
			EN: "Vacancy responds",
			RU: "Отклики",
		},
		"vacancy_respond": {
			EN: "Vacancy respond",
			RU: "Отклик",
		},
		"site_categories": {
			EN: "Site categories",
			RU: "Категории сайта",
		},
		"locations": {
			EN: "Locations",
			RU: "Локации",
		},
		"phone": {
			EN: "Phone",
			RU: "Телефон",
		},
		"vacancies": {
			EN: "Vacancies",
			RU: "Вакансии",
		},
		"resumes": {
			EN: "Resumes",
			RU: "Резюме",
		},
		"offerings": {
			EN: "Offerings",
			RU: "Услуги",
		},
		"realty": {
			EN: "Realty",
			RU: "Недвижимость",
		},
		"id_parameter_is_required": {
			EN: "ID parameter is required",
			RU: "ID обязателен",
		},
		"user_id_is_required": {
			EN: "User ID is required",
			RU: "user ID обязателен",
		},
		"username_required": {
			EN: "Username is required",
			RU: "Имя пользователя обязательно",
		},
		"username_alphanum": {
			EN: "Username must contain only letters and numbers",
			RU: "Имя пользователя должно содержать только буквы и цифры",
		},
		"username_min_length": {
			EN: "Username must be at least 1 character",
			RU: "Имя пользователя должно быть не менее 1 символа",
		},
		"username_max_length": {
			EN: "Username must be at most 25 characters",
			RU: "Имя пользователя должно быть не более 25 символов",
		},
		"email_required": {
			EN: "Email is required",
			RU: "Email обязателен",
		},
		"email_invalid": {
			EN: "Email is invalid",
			RU: "Email некорректный",
		},
		"email_min_length": {
			EN: "Email must be at least 1 character",
			RU: "Email должен быть не менее 1 символа",
		},
		"email_max_length": {
			EN: "Email must be at most 100 characters",
			RU: "Email должен быть не более 100 символов",
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
			EN: "Password must be at most 32 characters",
			RU: "Пароль должен быть не более 32 символов",
		},
		"password_confirm_required": {
			EN: "Password confirmation is required",
			RU: "Подтверждение пароля обязательно",
		},
		"password_confirm_mismatch": {
			EN: "Password confirmation does not match password",
			RU: "Подтверждение пароля не совпадает с паролем",
		},
		"quick_announcement_description_required": {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		"quick_announcement_user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязательно",
		},
		"quick_announcement_location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязательно",
		},
		"quick_announcement_site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязательно",
		},
		"quick_announcement_site_category_id2_required": {
			EN: "Quick announcement site category ID is required",
			RU: "ID категории быстрого объявления обязательно",
		},
		"quick_announcement_respond_id_required": {
			EN: "Quick announcement ID is required",
			RU: "ID быстрого объявления обязательно",
		},
		"quick_announcement_respond_user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязательно",
		},
		"messenger_name_required": {
			EN: "Messenger name is required",
			RU: "Имя мессенджера обязательно",
		},
		"messenger_group_id_required": {
			EN: "Group ID is required",
			RU: "ID группы обязательно",
		},
		"messenger_location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязательно",
		},
		"messenger_site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязательно",
		},
		"realty_site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязательно",
		},
		"realty_user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязательно",
		},
		"realty_location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязательно",
		},
		"realty_title_required": {
			EN: "Title is required",
			RU: "Заголовок обязателен",
		},
		"realty_description_required": {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		"vacancy_site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязательно",
		},
		"vacancy_title_required": {
			EN: "Title is required",
			RU: "Заголовок обязателен",
		},
		"vacancy_description_required": {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		"vacancy_currency_id_required": {
			EN: "Currency ID is required",
			RU: "ID валюты обязательно",
		},
		"vacancy_user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязательно",
		},
		"vacancy_location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязательно",
		},
		"offering_site_category_id_required": {
			EN: "Site category ID is required",
			RU: "ID категории сайта обязательно",
		},
		"offering_title_required": {
			EN: "Title is required",
			RU: "Заголовок обязателен",
		},
		"offering_description_required": {
			EN: "Description is required",
			RU: "Описание обязательно",
		},
		"offering_user_id_required": {
			EN: "User ID is required",
			RU: "ID пользователя обязательно",
		},
		"offering_location_id_required": {
			EN: "Location ID is required",
			RU: "ID локации обязательно",
		},
		"subscription_subscriber_id_required": {
			EN: "Subscriber ID is required",
			RU: "ID подписчика обязательно",
		},
		"subscription_target_user_id_required": {
			EN: "Target user ID is required",
			RU: "ID целевого пользователя обязательно",
		},
	}

}

// Translate возвращает перевод по ключу и языку
func Translate(key string, lang Language) string {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := translations[key]; ok {
		if tr, ok := val[lang]; ok {
			return tr
		}
		if tr, ok := val[EN]; ok {
			return tr
		}
	}
	return key
}
