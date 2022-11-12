## МИКРОСЕРВИС ДЛЯ РАБОТЫ С БАНАЛАНСОМ ПОЛЬЗОВАТЕЛЯ

Микросервис разработан в рамках тестового задания Avito Backend Internship.

## СОДЕРЖАНИЕ

[Инструкция по запуску](#instruc)

[Примеры запросов и ответов](#examples)

[Вопросы по ТЗ](#ques)

## Инструкция по запуску <a name="instruc"></a>

1. `$ git clone https://github.com/cubaki5/balance_manager_microservice`

2. `$ docker compose build`

3. `$ docker compose up`

Приложение запущено и готово к работе!

## Примеры запросов и ответов <a name="examples"></a>

Ознакомиться со Swagger 2.0 API документаций можно [здесь](http://localhost:1323/swagger/index.html) при запущенном приложении.

Примеры curl запросов:

- начисление средств на определённый аккаунт

  `$ curl --location --request POST 'http://localhost:1323/fund/accrual' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3,
  "income":1500
  }'`

- резервирование средств на аккаунте

    `$ curl --location --request POST 'http://localhost:1323/fund/reservation' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3,
  "service_id":666,
  "order_id":83,
  "cost":300
  }'`

- признание выручки

  `$ curl --location --request POST 'http://localhost:1323/fund/payment_acceptance' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3,
  "service_id":666,
  "order_id":83,
  "cost":300
  }'`

- разрезервирование средств

  `$ curl --location --request POST 'http://localhost:1323/fund/payment_rejection' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3,
  "service_id":666,
  "order_id":83,
  "cost":300
  }'`

- информация о балансе пользователя

  `$ curl --location --request GET 'http://localhost:1323/fund/balance' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3
  }'`

    Response Body:

    `{
        "balance":160
    }`

- отчёт для бухгалтерии

    `$ curl --location --request GET 'http://localhost:1323/report/accounting' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "year":2022,
  "month":11
  }'`

  Response Body:

  ``` 
  { 
  11;15900
  12;13000
   }
- история транзакций с возможностью сортировки по дате и сумме и предусмотренной пагинацией

    `$curl --location --request GET 'http://localhost:1323/report/transactions_history?sortDate=0&sortSum=1&page=1&limit=100' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id":3
  }'`

  Response Body:

    ``` 
    { 
    начисление;источник пополнения;2022-11-12T19:28:27Z;1500
    оплата;оплата заказа №1 по покупке 33;2022-12-12T19:28:27Z;600
    возврат;возврат стоимости заказа №2 по покупке 43;2022-12-14T19:28:27Z;250
     }

## Вопросы по ТЗ  <a name="ques"></a>
1. Какая нагрузка на микросервис планируется?
    - *На данный момент все запросы к базе данных обрабатываются в транзакциях с уровнем изоляции REPEATABLE READ. Однако уровень изоляции для операций по сохранению истории можно понизить, что позволит увеличить скорость работы.*
2. Нужно ли учитывать копейки при работе с денежными средствами?
   - *Денежные средства считаются целочисленным типом.*
3. В фунте ТЗ "Задача" говорится также о реализации "перевода средств от пользователя к пользователю", нужно ли его реализовать?
    - *В описании задание об этой задаче не говорилось, поэтому данная операция не предусмотрена в созданном микросервисе, однако может быть добавлена.*
4. Нужно ли в истории транзакций сохранять записи по резервированию и разрезервированию средств?
    - *В истории транзакций сохраняется эта информация.*
5. В примере отчёта для бухгалтерии указаны названия услуг, однако не сказано, откуда брать названия услуг.
    - *Вместо названия услуг микросервис в отчёте указывает ID услуги.*
6. В примере отчёта для бухгалтерии нет шапки таблицы, делать ли её?
    - *Микросервис возвращает отчёт без шапки, как в примере.*
7. В дополнительном задании 1 на выходе необходимо отдавать "ссылку на CSV файл", однако неясно, где хранить этот файл, какое у него должно быть название и тп.
    - *Из-за неопределённостей, при решении доп задания 1 было решено давать не ссылку на csv файл, а данные, в формате csv.*
8. В дополнительном задании 2 необходимо формировать список с "комментариями откуда и зачем были начислены/списаны средства", однако неясно, откуда брать конкретные названия услуг.
    - *При формировании списка по заданию 2 микросервис использует ID услуги.*
9. Также в доп задании 2 необходимо предусмотреть пагинацию. Появился вопрос, что выводить в ответе, когда по введённым параметрам на странице нет записей об истории (история короче), пустую страницу или записи выше?
    - *В такой ситуации микросервис возвращает пустое тело.*
10. По формулировке задания 2 необходимо предусмотреть сортировку "по дате и сумме". Как должна выглядеть сортировка и по дате, и по сумме?
    - *В параметрах данного запроса есть возможность поставить сортировку и по дате, и по сумме. Однако, механизм сортировки отдан на откуп базе данных.*
11. Можно ли доверять сервисам, которые используют данный микросервис?
    - *Так как микросервис работает с таким важными вещами, как баланс пользователя, нужно добавить валидацию данных, пришедших из вне.*