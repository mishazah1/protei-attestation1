package Использование_контекста

import (
	"context"
	"time"
)

// Управление жизненным циклом
func FetchData(ctx context.Context) (string, error) {
	// Создаем контекст с таймаутом на 2 секунды
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Освобождает ресурсы, если функция завершится быстрее таймаута

	resChan := make(chan string, 1)

	go func() {
		// Имитация долгой работы
		time.Sleep(3 * time.Second)
		resChan <- "data"
	}()

	select {
	case <-ctx.Done():
		// Контекст отменен или вышел таймаут
		return "", ctx.Err() // Вернет context.DeadlineExceeded или context.Canceled
	case res := <-resChan:
		return res, nil
	}
}

// ---------------------------------------------------------------------------------------------------------
// Передача значений
type contextKey string // или даже custom type для полной безопасности

const UserIDKey contextKey = "userID"

// Функция-хелпер для установки значения
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// Функция-хелпер для извлечения значения
func UserIDFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(UserIDKey).(string)
	return val, ok
}

// ---------------------------------------------------------------------------------------------------------
// Каскадная отмена
func ProcessRequest(ctx context.Context) {
	// Создаем контекст с отменой для всех подзадач
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go task1(ctx)
	go task2(ctx)

	// Если task1 обнаружил критическую ошибку, он может вызвать cancel(),
	// и task2 также будет остановлен автоматически через ctx.Done()
}
