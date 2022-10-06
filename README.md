# Fake OpenRTB DSP Server
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/RapidCodeLab/fakedsp/Deploy%20Docker%20image%20to%20Docker%20Hub?style=flat-square)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/rapidcodelab/fakedsp?style=flat-square)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/rapidcodelab/fakedsp/latest?style=flat-square)
[![GitHub license](https://img.shields.io/github/license/RapidCodeLab/fakedsp?style=flat-square)](https://github.com/RapidCodeLab/fakedsp/blob/main/LICENSE)
![GitHub top language](https://img.shields.io/github/languages/top/RapidCodeLab/fakedsp?style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RapidCodeLab/fakedsp?style=flat-square)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/RapidCodeLab/fakedsp/main?style=flat-square)

## OpenRTB 2.6 & OpenRTB Native Ads Specification 1.2

#### Фейковая DSP(Demand Side Platform) для валидации OpenRTB BidRequest запросов форматов Native, Banner, Video, Audio с отдачей валидных OpenRTB BidResponse ответов. 



#### Основные особенности и возможности.


* Валидируются только объекты и значения BidRequest указаные как required в спецификации. 

* Полная поддержка спецификации OpenRTB 2.6 & OpenRTB Native Ads Specification 1.2, включая множетственные объекты Imp в одном BidRequest запросе, а так же соместный запрос Native, Banner, Video, Audio объектов в одном объекте Imp ( п. 3.2.4 Спецификации OpenRTB 2.6 ). 

* На каждый объект Native, Banner, Video, Audio объектов в одном объекте Imp фейковая DSP отвечает BidResponse объектом с  необхдимым количеством Bid объектов с указанием impid & mtype в соответствии с спецификацией.

* В ответе BidResponse поддерживается множество SeatBid объектов с множеством Bid объектов, в соответствии с спецификацией.


#### Запуск в docker контейнере

```bash
docker run --name fakedsp -p 8080:8080 -v /path/to/local/ads.json:/ads.json -e ADS_DATABASE_PATH='./ads.json' rapidcodelab/fakedsp
```


#### Пример запроса


Запросы прнинимаютя на эндпоинте POST http://127.0.0.1:8080/openrtb


В одном BidRequest запросе два объекта Imp. 
Первый Imp содержит Native и Banner объекты.
Второй Imp содержит Video объект.  


#### OpenRTB BidRequest ()
```json
{
  "id": "1",
  "imp": [
    {
      "id":"1",
      "native":{
        "request":"{\"ver\":\"1.2\",\"layout\":1,\"adunit\":2,\"plcmtcnt\":4,\"plcmttype\":4,\"assets\":[{\"id\":1,\"required\":1,\"title\":{\"len\":75}},{\"id\":2,\"required\":1,\"img\":{\"wmin\":492,\"hmin\":328,\"type\":3,\"mimes\":[\"image/jpeg\",\"image/jpg\",\"image/png\",\"image/gif\"]}},{\"id\":4,\"required\":0,\"data\":{\"type\":6}},{\"id\":5,\"required\":0,\"data\":{\"type\":7}},{\"id\":6,\"required\":0,\"data\":{\"type\":1,\"len\":20}}]}"
      },
      "banner":{}
    },{
      "id": "2",
      "video": {
        "mimes": [
          "video/mp4"
        ],
        "minduration": 5,
        "maxduration": 30,
        "protocols": [
          2,
          3
        ]
      },
      "bidfloor": 0.5,
      "bidfloorcur": "USD",
      "ext":{}
    }
  ]
}
```

#### OpenRTB BidResponse ()
```json
{
  "id": "1",
  "seatbid": [
    {
      "bid": [
        {
          "id": "036881a7-d7cb-47a8-b07e-c9e6110fa4a0",
          "impid": "1",
          "price": 1.1046602879796197,
          "adm": "<a href=\"https://yandex.ru\"><img srec=\"https://banners.rapidcodelab.repl.co/banners/1.jpg\"/></a>",
          "mtype": 1
        },
        {
          "id": "80b1b7f4-d1be-4b42-94c7-cd706df46155",
          "impid": "1",
          "price": 1.4405090880450124,
          "adm": "<div><a href=\"https://yandex.ru\"><img src=\"https://banners.rapidcodelab.repl.co/banners/1.jpg\"/><br>Wordpress Hosting</a><br>Cheap wordpress hosting at turhost.com</div>",
          "mtype": 4
        },
        {
          "id": "de055180-9335-47b8-b6f3-92bd67531a2a",
          "impid": "2",
          "price": 1.1645600532184903,
          "adm": "<VAST version=\"3.0\"><Ad id=\"123\"><InLine><AdSystem><![CDATA[DSP]]></AdSystem><AdTitle><![CDATA[adTitle]]></AdTitle><Creatives><Creative><Linear><TrackingEvents></TrackingEvents><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" width=\"0\" height=\"0\"><![CDATA[https://banners.rapidcodelab.repl.co/media/002.mp4]]></MediaFile></MediaFiles><VideoClicks><ClickThrough id=\"1\"><![CDATA[https://yandex.ru]]></ClickThrough></VideoClicks></Linear></Creative></Creatives></InLine></Ad></VAST>",
          "mtype": 2
        }
      ],
      "seat": "creatives.com"
    }
  ]
}
```