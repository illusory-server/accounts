package repository

/*
Паттерны для достижения транзакицй
1. Использование аггретгатов как представляение что 2 разных данных имеют единый транзакционный случай.
Работает только с 1 репозиторием, не умеет работать с 2 и более по причине что аггрегат одна транзакционнность, 2 репозиториев это уже 2 разных
2. Если же агрегат невозможен с виду что данные отвечают за разные вещи, применить транзакцию на уровне БД.
Есть проблема что данный прием не работает с разными хранилищами или с взаимодействием с другим клиентом, но умеет работать с 2 и более репозиториев
3. Паттерн Unit Of Work, объект с необходимыми зависимостями и callback функцией которая отрабатывает в Do методе.
Приемущество в том что работает с множеством репозиториев а так же может работать с разными БД и даже с клиентами в виду своей реализаций как интерфейс зависимость
4. Паттерн Saga
*/
