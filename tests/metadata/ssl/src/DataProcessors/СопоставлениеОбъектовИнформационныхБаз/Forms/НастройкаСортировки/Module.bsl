///////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2021, ООО 1С-Софт
// Все права защищены. Эта программа и сопроводительные материалы предоставляются 
// в соответствии с условиями лицензии Attribution 4.0 International (CC BY 4.0)
// Текст лицензии доступен по ссылке:
// https://creativecommons.org/licenses/by/4.0/legalcode
///////////////////////////////////////////////////////////////////////////////////////////////////////

#Область ОбработчикиСобытийФормы

&НаСервере
Процедура ПриСозданииНаСервере(Отказ, СтандартнаяОбработка)
	
	ТаблицаСортировки.Загрузить(Параметры.ТаблицаСортировки.Выгрузить());
	
КонецПроцедуры

#КонецОбласти

#Область ОбработчикиКомандФормы

&НаКлиенте
Процедура Применить(Команда)
	
	ОповеститьОВыборе(ТаблицаСортировки);
	
КонецПроцедуры

&НаКлиенте
Процедура Отменить(Команда)
	
	ОповеститьОВыборе(Неопределено);
	
КонецПроцедуры

#КонецОбласти
