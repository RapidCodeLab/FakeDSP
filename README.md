# Fake OpenRTB DSP Server
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/RapidCodeLab/fakedsp/Deploy%20Docker%20image%20to%20Docker%20Hub?style=flat-square)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/rapidcodelab/fakedsp?style=flat-square)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/rapidcodelab/fakedsp/latest?style=flat-square)
[![GitHub license](https://img.shields.io/github/license/RapidCodeLab/fakedsp?style=flat-square)](https://github.com/RapidCodeLab/fakedsp/blob/main/LICENSE)
![GitHub top language](https://img.shields.io/github/languages/top/RapidCodeLab/fakedsp?style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RapidCodeLab/fakedsp?style=flat-square)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/RapidCodeLab/fakedsp/main?style=flat-square)

## OpenRTB 2.6 & OpenRTB Native Ads Specification 1.2

### Фейковая DSP(Demand Side Platform) для валидации OpenRTB BidRequest запросов форматов Native, Banner, Video, Audio с отдачей валидных OpenRTB BidResponse ответов. 



### Основные особенности и возможности.


* Валидируются только объекты и значения BidRequest указаные как required в спецификации. 

* Полная поддержка спецификации OpenRTB 2.6 & OpenRTB Native Ads Specification 1.2, включая множетственные объекты Imp в одном BidRequest запросе, а так же соместный запрос Native, Banner, Video, Audio объектов в одном объекте Imp ( п. 3.2.4 Спецификации OpenRTB 2.6 ). 

* На каждый объект Native, Banner, Video, Audio объектов в одном объекте Imp фейковая DSP отвечает BidResponse объектом с  необхдимым количеством Bid объектов с указанием impid & mtype в соответствии с спецификацией.

* В ответе BidResponse поддерживается множество SeatBid объектов с множеством Bid объектов, в соответствии с спецификацией.