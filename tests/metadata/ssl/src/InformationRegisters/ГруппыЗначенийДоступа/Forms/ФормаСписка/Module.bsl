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

	УстановитьУсловноеОформление();
	
	ТолькоПросмотр = Истина;
	
КонецПроцедуры

#КонецОбласти

#Область ОбработчикиКомандФормы

&НаКлиенте
Процедура ВключитьВозможностьРедактирования(Команда)
	
	ТолькоПросмотр = Ложь;
	
КонецПроцедуры

&НаКлиенте
Процедура ОбновитьДанныеРегистра(Команда)
	
	ЕстьИзменения = Ложь;
	
	ОбновитьДанныеРегистраНаСервере(ЕстьИзменения);
	
	Если ЕстьИзменения Тогда
		Текст = НСтр("ru = 'Обновление выполнено успешно.'");
	Иначе
		Текст = НСтр("ru = 'Обновление не требуется.'");
	КонецЕсли;
	
	ПоказатьПредупреждение(, Текст);
	
КонецПроцедуры

#КонецОбласти

#Область СлужебныеПроцедурыИФункции

&НаСервере
Процедура УстановитьУсловноеОформление()
	
	Список.КомпоновщикНастроек.Настройки.УсловноеОформление.Элементы.Очистить();
	
	ОформитьГруппуДанных(0, НСтр("ru = 'Стандартные значения доступа'"));
	ОформитьГруппуДанных(1, НСтр("ru = 'Обычные/внешние пользователи'"));
	ОформитьГруппуДанных(2, НСтр("ru = 'Обычные/внешние группы пользователей'"));
	ОформитьГруппуДанных(3, НСтр("ru = 'Группы исполнителей'"));
	ОформитьГруппуДанных(4, НСтр("ru = 'Объекты авторизации'"));
	
КонецПроцедуры

&НаСервере
Процедура ОформитьГруппуДанных(ГруппаДанных, Текст)
	
	ЭлементОформления = Список.КомпоновщикНастроек.Настройки.УсловноеОформление.Элементы.Добавить();
	ЭлементОформления.РежимОтображения = РежимОтображенияЭлементаНастройкиКомпоновкиДанных.Недоступный;
	
	ЭлементПоля = ЭлементОформления.Поля.Элементы.Добавить();
	ЭлементПоля.Поле = Новый ПолеКомпоновкиДанных("ГруппаДанных");
	
	ЭлементОтбора = ЭлементОформления.Отбор.Элементы.Добавить(Тип("ЭлементОтбораКомпоновкиДанных"));
	ЭлементОтбора.ЛевоеЗначение = Новый ПолеКомпоновкиДанных("ГруппаДанных");
	ЭлементОтбора.ВидСравнения = ВидСравненияКомпоновкиДанных.Равно;
	ЭлементОтбора.ПравоеЗначение = ГруппаДанных;
	
	ЭлементОформления.Оформление.УстановитьЗначениеПараметра("Текст", Текст);
	
КонецПроцедуры

&НаСервере
Процедура ОбновитьДанныеРегистраНаСервере(ЕстьИзменения)
	
	УстановитьПривилегированныйРежим(Истина);
	
	РегистрыСведений.ГруппыЗначенийДоступа.ОбновитьДанныеРегистра(ЕстьИзменения);
	
	Элементы.Список.Обновить();
	
КонецПроцедуры

#КонецОбласти
