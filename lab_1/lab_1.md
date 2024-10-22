# Лабораторная работа №1

<h2>Цель лабораторной работы №1</h2>

Цель данной работы – ознакомится с применением протокола HTTP на практике, в реальных системах. Каждый из рассмотренных типов запросов предлагается отправить на несколько известных интернет-сервисов. Впрочем, сервисы указаны лишь как примеры и при желании вы можете выбрать другие (социальные сети, почта, облака, новостные сайты и т.д.).  
С помощью специального ПО (Postman, либо многочисленные аналоги, например, Restlet Clent - расширение для Chrome) вручную отправить следующие запросы и ответить на предлагаемые вопросы.



<h2>Задание</h2>
<p>Необходимо провести ознакомительные работы с протоколом HTTP.

<h3>1.2.1. Ознакомиться с методом OPTIONS.</h3>

Данный метод предназначен для получения информации о сервере: проверка работоспособности, тестирования на возможности.</p>

1) Отправим запрос на https://mail.ru . 

Ответ:

![](img/option/mail_https.png)

В данном ответе был получен ответ об ошибке 400 (Bad Request). 

Данная ошибка говорит о том, что сервер не смог обработать данный запрос. Эта ошибка появляется, когда было некорректно указан запрос. 


2) Оправим запрос на https://ya.ru

Ответ:

![](img/option/ya_https.png)

В данном запросе ответом была получена 200. 

3) Отправим запрос на https://vk.com

![](img/option/vk.png)

В данном запросе был получен ответ 418.
Данный код ошибки был добавлен первого апреля в 1998 году. Он говорит о том, что сервер не может приготовить кофе, так как он чайник.

4) Отправим запрос на https://rambler.ru

![](img/option/rambler.png)

Данный запрос выдал код 200. 

5) Отпрвим запрос на hyyps://ok.ru

![](img/option/ok.png)

Данный запрос выдал код 200, а так же в ответе пы получили корректный ответ на данный метод. 

В haeders Allow были получены методы GET, HEAD, POST, TRACE, OPTIONS.

<p>Для чего используется запрос OPTIONS? Какие коды ответов приходят при этом запросе? Какие сайты правильно обработали запрос и вернули ожидаемые данные?</p>

Данный метод используется для получениия информации о состоянии сервера, а так же его возможностей. 
В данном разделе были получены ответы 200, 400, 418.
<ul>
    <li>200 - код, который говорит о том, что запрос прошел корректно
    <li>400 - код, который говорит о том, что был отправлен некоректный запрос
    <li>418 - код, который говорит о том, сервер отаправил шутку в ответ. В этом ответе нет никакой валидной информации.
</ul>

Из предложенного списка ответов все сайты кроме одного отправили коректный ответ. 

Только https://ok.ru отправил коректный ответ, все остальные серверы отправили некоректный ответ. 

<h3>1.2.2 Ознакомиться с методом HEAD.</h3>

1) Отправим запрос на https://vk.com

![](img/head/vk.png)

В данном запросе был получен ответ 418 (я не чайник)

2) Отправим запрос на https://msn.com

![](img/head/msn.png)

В данном запросе был получен ответ 404 (страница не нейдена) Это говорит о том, что во время запроса не был найден корректный ответ.

3) Отправим запрос на https://apple.com

![](img/head/apple.png)

Данный запрос сработал корректно. Было получен ответ 200. Из интересного что я бы подметил: заголовок connection : keep-alive. Я бы предположил, что это можно считать как раз разрешением на следующий запрос. 


<p>Метод HEAD выполнает вспомогательные функции.
Он необходи для того что бы получать информацию от сервера или проверять его работу или валидность запроса.</p>

Отличие HEAD от GET заключается в том, что GET может передвать что-то в теле запроса, а вот HEAD не будет отправлять информацию в теле запроса.

Я бы сказал, что корректный ответ был получен только при запросе к серверу apple.


<h3>1.2.3 Ознакомиться с методами POST и GET.</h3>

•	GET — метод для получения данных с сервера. Он передаётся с URL, поэтому виден в адресной строке браузера или истории посещений сайтов. Применяется в фильтрах поисковиков и онлайн-магазинов.

•	POST — метод отправки данных на сервер, например, после заполнения формы регистрации или авторизации на сайте.


<h4>GET метод:</h4>

1) Оправим запрос на https://yandex.ru

Headers:
![](img/get/yandex_1.png)
body:
![](img/get/yandex_2.png)
В теле ответа видно, что сервер отправил ответом html страницу. 

2) Отправим запрос на https://google.com
Headers:
![](img/get/google_1.png)
body:
![](img/get/google_2.png)
В данном запросе так же была получена страница html.

3) Отправим запрос на https://apple.com
Headers:
![](img/get/apple_1.png)
body:
![](img/get/apple_2.png)
В данном запросе так же была получена страница html.

<h4>POST метод:</h4>

1) Отправим запрос на https://yandex.ru

body request:

```
{
  "title": "foo",
  "body": "bar",
  "userId": 1
}
```

Headers response:
![](img/post/yandex_1.png)
body response:
![](img/post/yandex_2.png)

В данном ответе видно, что при смене метода был и отправлен сервером другая html страница. 

2) Отправим запрос на https://google.com

body request:

```
{
  "title": "foo",
  "body": "bar",
  "userId": 1
}
```

Headers response:
![](img/post/google_1.png)
body response:
![](img/post/google_2.png)

В данном запросе была получен код ошибки 405.
Данная ошибка говорит о том, что данный метод запрещен по данному url адресу. было получен html страница.


3) Отправим запрос на https://apple.com

body request:

```
{
  "title": "foo",
  "body": "bar",
  "userId": 1
}
```

Headers response:
![](img/post/apple_1.png)
body response:
![](img/post/apple_2.png)

При данном запросе был получен код 200. 
И сервер в ответ отправляет html страницу.

<h3>1.3. Работа с API VK</h3>

Был получен токен доступа.

1) Отправим запрос на https://api.vk.com/method/users.get?user_ids=601199047&fields=photo_200&access_token=vk1.a.DJmVXkrWrDQ3z91gqDVGEIQE-bgLpMALr8vrmcfsuVd-YY4EPGKNeNBz96QjuOI2Fai8o-VNvJKFk5BrK50rDdzJEIvDohnc2WFIi7nttPZPfrKUbJBrrboWSdR3VBtsadEGAcgTEeS6wswMT6WbGzDwg4qz0nxEHmx6vz4C6GxYLYxzewjPFTCH6nuUr9PTbcxXErtSK_qm-_LRLK84Lw&v=5.199 HTTP/1.1


Получение фото:
![](img/vk/1.png)

Фото, которое было получено по ссылке:
![](img/vk/2.png)

<h3>1.3.2.1. Получите список всех факультетов МГТУ им. Н.Э.Баумана.</h3>

1) Отправим запрос на https://api.vk.com/method/database.getFaculties?university_id=250&access_token=vk1.a.WklpDcfM_VEReIaGV6xc0RAU6mslLOZUTpQ-WcfJdvu8RQxZhWgMFKsfNF-RI4LuvSqZHVMyR5-D6fUEtQlq4poVcX1G-A2EN_I16pj0zdu31qiVSmmjhEspbj_S0xEFOAp-lI36gecuFLF5lKqwmioSIQ8zLrMgtbZwDWCHyN2tY4aCGBlJZk9Mk114rVkayrduve9Fpi7-rVySLCewEQ&v=5.199 HTTP/1.1

Получение фото:
![](img/vk/3.png)


<h3>1.3.3.1. Отправить запись на стену.</h3>

1) Отправим запрос на https://api.vk.com/method/wall.post?message=testing&owner_id=601199047&access_token=vk1.a.WklpDcfM_VEReIaGV6xc0RAU6mslLOZUTpQ-WcfJdvu8RQxZhWgMFKsfNF-RI4LuvSqZHVMyR5-D6fUEtQlq4poVcX1G-A2EN_I16pj0zdu31qiVSmmjhEspbj_S0xEFOAp-lI36gecuFLF5lKqwmioSIQ8zLrMgtbZwDWCHyN2tY4aCGBlJZk9Mk114rVkayrduve9Fpi7-rVySLCewEQ&expires_in=86400&v=5.199 HTTP/1.1

Ответ на отправку запроса на запись на стену:
![](img/vk/4.png)

Результат запроса:
![](img/vk/5.png)



















