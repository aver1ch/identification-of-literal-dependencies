# 🔍 Text Plagiarism Detection Service (Winnowing Algorithm in Go)

Этот сервис предназначен для **выявления текстовых заимствований** между `.txt` документами с помощью алгоритма **Winnowing Fingerprinting**, реализованного на языке Go.

## 📦 Возможности

- Быстрое сравнение текстов на уровне фрагментов (символов или слов)
- Указание точных **источников заимствований**
- Поддержка большого числа документов
- Гибко настраиваемые параметры (длина n-граммы, окно и т.д.)
- Устойчивость к незначительным изменениям в тексте
- Генерация **отчёта** со списком совпадений и источников

---

## 🚀 Как работает

Алгоритм проходит следующие этапы:

1. **Очистка текста**: удаление пунктуации, перевод в нижний регистр
2. **Генерация n-грамм**: нарезка текста на фрагменты фиксированной длины
3. **Хэширование n-грамм**: создание цифровых отпечатков (hash)
4. **Winnowing**: выбор минимальных хэшей из скользящих окон
5. **Индексирование** всех документов из базы
6. **Сравнение нового текста** с индексом — выявление совпадений
7. **Формирование отчёта**: какие фрагменты и с какими источниками совпадают

---

## 🛠 Установка

1. Склонируйте репозиторий:

```bash
git clone https://github.com/yourusername/plagiarism-checker-go.git
cd plagiarism-checker-go
