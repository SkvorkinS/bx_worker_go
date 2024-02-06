package bx_worker_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Shmitrix struct {
	Webhook string
}

func (b Shmitrix) CrmDealList(filter map[string]interface{}, fields []string) (map[string]interface{}, error) {
	// Создаем объект с данными, которые хотим отправить
	data := map[string]interface{}{
		"filter": filter,
		"select": fields,
	}

	// Конвертируем объект в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%v/crm.deal.list", b.Webhook)

	// Создаем новый HTTP-запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Устанавливаем заголовок Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b Shmitrix) CrmDealUpdate(id string, fields map[string]interface{}) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ID":     id,
		"FIELDS": fields,
		"params": map[string]string{
			"REGISTER_SONET_EVENT": "N",
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Замените URL на ваш URL-адрес сервера
	url := fmt.Sprintf("%v/crm.deal.update", b.Webhook)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return result, nil
}

func (b Shmitrix) CrmDealAdd(fields map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/crm.deal.add.json", b.Webhook)

	// Данные для запроса
	data := map[string]interface{}{
		"FIELDS": fields,
		"params": map[string]string{
			"REGISTER_SONET_EVENT": "N",
		},
	}

	// Преобразование данных в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Создание запроса
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверка, что тело ответа не nil
	if resp.Body == nil {
		return nil, errors.New("response body is nil")
	}

	// Чтение тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
