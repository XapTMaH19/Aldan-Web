package models

type Schema struct {
	NSI           []NSIEntity        `json:"NSI"`
	BloatedSets   BloatedSetsEntity  `json:"bloatedSets"`
	Charts        []interface{}      `json:"charts"`
	ConfigVersion int64              `json:"configVersion"`
	CurrentZoom   int64              `json:"currentZoom"`
	Floats        []FloatsEntity     `json:"floats"`
	Lines         []LinesEntity      `json:"lines"`
	MaxLineID     int64              `json:"maxLineID"`
	MaxNSIID      int64              `json:"maxNSIID"`
	MaxObjectID   int64              `json:"maxObjectID"`
	Objects       []ObjectsEntity    `json:"objects"`
	References    []ReferencesEntity `json:"references"`
	Result        string             `json:"result"`
	WorkingSet    string             `json:"workingSet"`
}

type NSIEntity struct {
	Class      string           `json:"class"`
	Id         int64            `json:"id"`
	ParentID   int64            `json:"parentID"`
	Properties PropertiesEntity `json:"properties"`
	X          int64            `json:"x"`
	Y          int64            `json:"y"`
	Z          int64            `json:"z"`
}

type PropertiesEntity struct {
	NormalField  NormalEntity  `json:"Данные регистрации"`
	NormalField1 NormalEntity1 `json:"Р"`
	NormalField2 NormalEntity2 `json:"Твоз"`
	NormalField3 NormalEntity3 `json:"Тгр"`
	Q            QEntity       `json:"Цена Qтг"`
	W            WEntity       `json:"Цена Wэл"`
}

type NormalEntity struct {
	Id   IdEntity   `json:"id"`
	Text TextEntity `json:"text"`
}

type IdEntity struct {
	Value string `json:"value"`
}

type TextEntity struct {
	Value string `json:"value"`
}

type NormalEntity1 struct {
	Value string `json:"value"`
}

type NormalEntity2 struct {
	Value string `json:"value"`
}

type NormalEntity3 struct {
	Value string `json:"value"`
}

type QEntity struct {
	Value string `json:"value"`
}

type WEntity struct {
	Value string `json:"value"`
}

type BloatedSetsEntity struct {
}

type FloatsEntity struct {
	CommonDP int64 `json:"Common.dP"`
}

type LinesEntity struct {
	EndID         int64          `json:"endID"`
	EndJunction   int64          `json:"endJunction"`
	Id            int64          `json:"id"`
	PipeID        int64          `json:"pipeID"`
	Points        []PointsEntity `json:"points"`
	StartID       int64          `json:"startID"`
	StartJunction int64          `json:"startJunction"`
}

type PointsEntity struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type ObjectsEntity struct {
	Class      string            `json:"class"`
	Id         int64             `json:"id"`
	ParentID   int64             `json:"parentID"`
	Properties PropertiesEntity4 `json:"properties"`
	ScrollX    int64             `json:"scrollX"`
	ScrollY    int64             `json:"scrollY"`
	X          int64             `json:"x"`
	Y          int64             `json:"y"`
	Z          int64             `json:"z"`
	ZValueMax  float64           `json:"zValueMax"`
	ZValueMin  int64             `json:"zValueMin"`
}

type PropertiesEntity4 struct {
	Design      DesignEntity   `json:"Design"`
	Inner       InnerEntity    `json:"Inner"`
	Tasks       TasksEntity    `json:"Tasks"`
	NormalField NormalEntity89 `json:"Данные регистрации"`
}

type DesignEntity struct {
	Bgcolor      BgcolorEntity    `json:"bgcolor"`
	ModelTime    ModelTimeEntity  `json:"modelTime"`
	WorkingSet   WorkingSetEntity `json:"workingSet"`
	NormalField  NormalEntity0    `json:"Перемычки"`
	NormalField1 NormalEntity5    `json:"Размеры"`
}

type BgcolorEntity struct {
	Value string `json:"value"`
}

type ModelTimeEntity struct {
	Value string `json:"value"`
}

type WorkingSetEntity struct {
	Value string `json:"value"`
}

type NormalEntity0 struct {
	Linecolor LinecolorEntity `json:"linecolor"`
	Linewidth LinewidthEntity `json:"linewidth"`
}

type LinecolorEntity struct {
	Value string `json:"value"`
}

type LinewidthEntity struct {
	Value string `json:"value"`
}

type NormalEntity5 struct {
	SceneHeight SceneHeightEntity `json:"sceneHeight"`
	SceneLeft   SceneLeftEntity   `json:"sceneLeft"`
	SceneTop    SceneTopEntity    `json:"sceneTop"`
	SceneWidth  SceneWidthEntity  `json:"sceneWidth"`
}

type SceneHeightEntity struct {
	Value string `json:"value"`
}

type SceneLeftEntity struct {
	Value string `json:"value"`
}

type SceneTopEntity struct {
	Value string `json:"value"`
}

type SceneWidthEntity struct {
	Value string `json:"value"`
}

type InnerEntity struct {
	LastError    LastErrorEntity  `json:"LastError"`
	Reserved     ReservedEntity   `json:"Reserved"`
	ScrollPosX   ScrollPosXEntity `json:"ScrollPosX"`
	ScrollPosY   ScrollPosYEntity `json:"ScrollPosY"`
	TabIndex     TabIndexEntity   `json:"TabIndex"`
	NormalField  NormalEntity6    `json:"Задачи"`
	NormalField1 NormalEntity7    `json:"Квазистационар"`
}

type LastErrorEntity struct {
	Value string `json:"value"`
}

type ReservedEntity struct {
	Value string `json:"value"`
}

type ScrollPosXEntity struct {
	Value string `json:"value"`
}

type ScrollPosYEntity struct {
	Value string `json:"value"`
}

type TabIndexEntity struct {
	Value string `json:"value"`
}

type NormalEntity6 struct {
	Value string `json:"value"`
}

type NormalEntity7 struct {
	Value string `json:"value"`
}

type TasksEntity struct {
	Task         TaskEntity         `json:"Дополнительные условия::Task"`
	Task1        TaskEntity8        `json:"Использовать динамику::Task"`
	Task2        TaskEntity9        `json:"Использовать индикаторные диаграммы ГДИ::Task"`
	Task3        TaskEntity10       `json:"Контроль ограничений::Task"`
	Task4        TaskEntity15       `json:"Максимальное количество итераций::Task"`
	Task5        TaskEntity16       `json:"Оценочные веса::Task"`
	Task6        TaskEntity31       `json:"Параметры адаптации модели::Task"`
	Task7        TaskEntity45       `json:"Параметры модели::Task"`
	Task8        TaskEntity51       `json:"Параметры расчета::Task"`
	PTask        PTaskEntity        `json:"Погрешность по P::Task"`
	QTask        QTaskEntity62      `json:"Погрешность по Q::Task"`
	TTask        TTaskEntity63      `json:"Погрешность по T::Task"`
	Task9        TaskEntity64       `json:"Подзадачи::Task"`
	Task10       TaskEntity65       `json:"Потоковый режим::Task"`
	Task11       TaskEntity66       `json:"Проект::Task"`
	Task12       TaskEntity67       `json:"Проектирование::Task"`
	Task13       TaskEntity72       `json:"Сигнализация (чувствительность)::Task"`
	Task14       TaskEntity86       `json:"Таймаут расчетного модуля::Task"`
	QTask15      QTaskEntity87      `json:"Точность по Qузл::Task"`
	Task16       TaskEntity88       `json:"Формировать отчеты по слоям::Task"`
	Normal10Task Normal10TaskEntity `json:"Шифр ГПУ-10::Task"`
}

type TaskEntity struct {
	Value string `json:"value"`
}

type TaskEntity8 struct {
	Value string `json:"value"`
}

type TaskEntity9 struct {
	Value string `json:"value"`
}

type TaskEntity10 struct {
	QTask  QTaskEntity   `json:"Допустимый % превышения Qмак скважин::Task"`
	KTask  KTaskEntity   `json:"Контроль Kзагр (ГПА)::Task"`
	KTask1 KTaskEntity11 `json:"Контроль Kуд (ГПА)::Task"`
	QTask2 QTaskEntity12 `json:"Контроль Qмак (ГПА)::Task"`
	Task   TaskEntity13  `json:"Контроль факт. Рмакс ТС::Task"`
	Task3  TaskEntity14  `json:"Фиксировать % окрытия кранов РД::Task"`
}

type QTaskEntity struct {
	Value string `json:"value"`
}

type KTaskEntity struct {
	Value string `json:"value"`
}

type KTaskEntity11 struct {
	Value string `json:"value"`
}

type QTaskEntity12 struct {
	Value string `json:"value"`
}

type TaskEntity13 struct {
	Value string `json:"value"`
}

type TaskEntity14 struct {
	Value string `json:"value"`
}

type TaskEntity15 struct {
	Value string `json:"value"`
}

type TaskEntity16 struct {
	Task     TaskEntity17   `json:"Аварийные отключения ГПА::Task"`
	Task1    TaskEntity18   `json:"Аварийные отключения ГРС::Task"`
	Task2    TaskEntity19   `json:"Изменения оборотов ГПА::Task"`
	Task3    TaskEntity20   `json:"Изменения режимов байпасирования КЦ::Task"`
	PmaxTask PmaxTaskEntity `json:"Нарушения по Pmax::Task"`
	TmaxTask TmaxTaskEntity `json:"Нарушения по Tmax::Task"`
	Task4    TaskEntity21   `json:"Нарушения по мощности ГПА::Task"`
	Task5    TaskEntity22   `json:"Нарушения по помпажу ГПА::Task"`
	Task6    TaskEntity23   `json:"Недопоставленный газ::Task"`
	Task7    TaskEntity24   `json:"Переключения ГПА::Task"`
	Task8    TaskEntity25   `json:"Переключения кранов::Task"`
	Task9    TaskEntity26   `json:"Потерянный газ::Task"`
	Task10   TaskEntity27   `json:"Разрывы труб::Task"`
	Task11   TaskEntity28   `json:"Соблюдение плана поставки на границах ГТС::Task"`
	Task12   TaskEntity29   `json:"Топливный газ::Task"`
	Task13   TaskEntity30   `json:"Электроэнергия::Task"`
}

type TaskEntity17 struct {
	Value string `json:"value"`
}

type TaskEntity18 struct {
	Value string `json:"value"`
}

type TaskEntity19 struct {
	Value string `json:"value"`
}

type TaskEntity20 struct {
	Value string `json:"value"`
}

type PmaxTaskEntity struct {
	Value string `json:"value"`
}

type TmaxTaskEntity struct {
	Value string `json:"value"`
}

type TaskEntity21 struct {
	Value string `json:"value"`
}

type TaskEntity22 struct {
	Value string `json:"value"`
}

type TaskEntity23 struct {
	Value string `json:"value"`
}

type TaskEntity24 struct {
	Value string `json:"value"`
}

type TaskEntity25 struct {
	Value string `json:"value"`
}

type TaskEntity26 struct {
	Value string `json:"value"`
}

type TaskEntity27 struct {
	Value string `json:"value"`
}

type TaskEntity28 struct {
	Value string `json:"value"`
}

type TaskEntity29 struct {
	Value string `json:"value"`
}

type TaskEntity30 struct {
	Value string `json:"value"`
}

type TaskEntity31 struct {
	Task   TaskEntity32 `json:"Коэф. гидравл. сопрот. крана редуцирования::Task"`
	Task1  TaskEntity33 `json:"Коэф. гидравл. эффективности::Task"`
	Task2  TaskEntity34 `json:"Коэф. теплообмена::Task"`
	Task3  TaskEntity35 `json:"Нач. приближения Кад::Task"`
	Task4  TaskEntity36 `json:"Ограничения на Кад КЦ::Task"`
	Task5  TaskEntity37 `json:"Ограничения на Кад::Task"`
	Task6  TaskEntity38 `json:"Полный интерфейс::Task"`
	Task7  TaskEntity39 `json:"Поправка на АВО КЦ::Task"`
	Task8  TaskEntity40 `json:"Поправка на обороты ГПА КЦ::Task"`
	Task9  TaskEntity41 `json:"Поправка на политроп. КПД ГПА КЦ::Task"`
	Task10 TaskEntity42 `json:"Поправка на степень сжатия ГПА КЦ::Task"`
	Task11 TaskEntity43 `json:"Расчет поправок краевых параметров::Task"`
	Task12 TaskEntity44 `json:"Условный диаметр крана редуцирования::Task"`
}

type TaskEntity32 struct {
	Value string `json:"value"`
}

type TaskEntity33 struct {
	Value string `json:"value"`
}

type TaskEntity34 struct {
	Value string `json:"value"`
}

type TaskEntity35 struct {
	Value string `json:"value"`
}

type TaskEntity36 struct {
	Value string `json:"value"`
}

type TaskEntity37 struct {
	Value string `json:"value"`
}

type TaskEntity38 struct {
	Value string `json:"value"`
}

type TaskEntity39 struct {
	Value string `json:"value"`
}

type TaskEntity40 struct {
	Value string `json:"value"`
}

type TaskEntity41 struct {
	Value string `json:"value"`
}

type TaskEntity42 struct {
	Value string `json:"value"`
}

type TaskEntity43 struct {
	Value string `json:"value"`
}

type TaskEntity44 struct {
	Value string `json:"value"`
}

type TaskEntity45 struct {
	Task  TaskEntity46 `json:"Автоопределение ступеней ГПА и состояний кранов обвязки::Task"`
	Task1 TaskEntity47 `json:"Автоотключение цеха::Task"`
	Task2 TaskEntity48 `json:"Запретить квазистационар::Task"`
	Task3 TaskEntity49 `json:"Игнорировать висячие узлы::Task"`
	Task4 TaskEntity50 `json:"Игнорировать трубы обвязки КЦ::Task"`
}

type TaskEntity46 struct {
	Value string `json:"value"`
}

type TaskEntity47 struct {
	Value string `json:"value"`
}

type TaskEntity48 struct {
	Value string `json:"value"`
}

type TaskEntity49 struct {
	Value string `json:"value"`
}

type TaskEntity50 struct {
	Value string `json:"value"`
}

type TaskEntity51 struct {
	NNTask NNTaskEntity  `json:"Nрасп/Nном <=::Task"`
	Task   TaskEntity52  `json:"Группировать кусты ГРС::Task"`
	KTask  KTaskEntity53 `json:"Использование Kад::Task"`
	Task1  TaskEntity54  `json:"Критерий оптимизации КЦ::Task"`
	Task2  TaskEntity55  `json:"Макс. точек разбиения::Task"`
	Task3  TaskEntity56  `json:"Метод расчета топливного газа::Task"`
	Task4  TaskEntity57  `json:"Мин. точек разбиения::Task"`
	Task5  TaskEntity58  `json:"Нач. приближения Р в узлах::Task"`
	TTask  TTaskEntity   `json:"Поправка на Tвоз::Task"`
	Task6  TaskEntity59  `json:"Расчет АВО::Task"`
	Task7  TaskEntity60  `json:"Свертка труб::Task"`
	Task8  TaskEntity61  `json:"Формулы ОНТП::Task"`
}

type NNTaskEntity struct {
	Value string `json:"value"`
}

type TaskEntity52 struct {
	Value string `json:"value"`
}

type KTaskEntity53 struct {
	Value string `json:"value"`
}

type TaskEntity54 struct {
	Value string `json:"value"`
}

type TaskEntity55 struct {
	Value string `json:"value"`
}

type TaskEntity56 struct {
	Value string `json:"value"`
}

type TaskEntity57 struct {
	Value string `json:"value"`
}

type TaskEntity58 struct {
	Value string `json:"value"`
}

type TTaskEntity struct {
	Value string `json:"value"`
}

type TaskEntity59 struct {
	Value string `json:"value"`
}

type TaskEntity60 struct {
	Value string `json:"value"`
}

type TaskEntity61 struct {
	Value string `json:"value"`
}

type PTaskEntity struct {
	Value string `json:"value"`
}

type QTaskEntity62 struct {
	Value string `json:"value"`
}

type TTaskEntity63 struct {
	Value string `json:"value"`
}

type TaskEntity64 struct {
	Value string `json:"value"`
}

type TaskEntity65 struct {
	Value string `json:"value"`
}

type TaskEntity66 struct {
	Value string `json:"value"`
}

type TaskEntity67 struct {
	DDTask DDTaskEntity `json:"Dнар реконст./ Dнар >::Task"`
	Task   TaskEntity68 `json:"Врезка кранов РД::Task"`
	Task1  TaskEntity69 `json:"Реконструкция, если год ввода в эксплуатацию трубопровода <::Task"`
	Task2  TaskEntity70 `json:"Свертка труб развития::Task"`
	Task3  TaskEntity71 `json:"Увеличение степени сжатия на КЦ::Task"`
}

type DDTaskEntity struct {
	Value string `json:"value"`
}

type TaskEntity68 struct {
	Value string `json:"value"`
}

type TaskEntity69 struct {
	Value string `json:"value"`
}

type TaskEntity70 struct {
	Value string `json:"value"`
}

type TaskEntity71 struct {
	Value string `json:"value"`
}

type TaskEntity72 struct {
	Task   TaskEntity73  `json:"Загрузка по мощности::Task"`
	Task1  TaskEntity74  `json:"Зона помпажа::Task"`
	Task2  TaskEntity75  `json:"Изменение процента открытия регулятора::Task"`
	Task3  TaskEntity76  `json:"Максимальная отношение скорость газа на выходе регулятора к скорости звука в газе::Task"`
	Task4  TaskEntity77  `json:"Максимальная скорость потока газа в фильтрах ГРС::Task"`
	Task5  TaskEntity78  `json:"Максимальная скорость потока газа::Task"`
	Task6  TaskEntity79  `json:"Максимальный % открытия регуляторов давления газа::Task"`
	Task7  TaskEntity80  `json:"Минимальная скорость потока газа в трубах::Task"`
	Task8  TaskEntity81  `json:"Режим вентиляции::Task"`
	ETask  ETaskEntity   `json:"Скорость изменения E::Task"`
	KTask  KTaskEntity82 `json:"Скорость изменения Kуд::Task"`
	Task9  TaskEntity83  `json:"Скорость изменения Кзаг::Task"`
	Task10 TaskEntity84  `json:"Скорость изменения давления::Task"`
	Task11 TaskEntity85  `json:"Чувствительность сравнительного отчета::Task"`
}

type TaskEntity73 struct {
	Value string `json:"value"`
}

type TaskEntity74 struct {
	Value string `json:"value"`
}

type TaskEntity75 struct {
	Value string `json:"value"`
}

type TaskEntity76 struct {
	Value string `json:"value"`
}

type TaskEntity77 struct {
	Value string `json:"value"`
}

type TaskEntity78 struct {
	Value string `json:"value"`
}

type TaskEntity79 struct {
	Value string `json:"value"`
}

type TaskEntity80 struct {
	Value string `json:"value"`
}

type TaskEntity81 struct {
	Value string `json:"value"`
}

type ETaskEntity struct {
	Value string `json:"value"`
}

type KTaskEntity82 struct {
	Value string `json:"value"`
}

type TaskEntity83 struct {
	Value string `json:"value"`
}

type TaskEntity84 struct {
	Value string `json:"value"`
}

type TaskEntity85 struct {
	Value string `json:"value"`
}

type TaskEntity86 struct {
	Value string `json:"value"`
}

type QTaskEntity87 struct {
	Value string `json:"value"`
}

type TaskEntity88 struct {
	Value string `json:"value"`
}

type Normal10TaskEntity struct {
	Value string `json:"value"`
}

type NormalEntity89 struct {
	Id           IdEntity       `json:"id"`
	Text         TextEntity     `json:"text"`
	NormalField  NormalEntity90 `json:"ГТО"`
	NormalField1 NormalEntity91 `json:"История изменений"`
	NormalField2 NormalEntity92 `json:"Описание схемы"`
	NormalField3 NormalEntity93 `json:"Пароль для полного доступа"`
	NormalField4 NormalEntity94 `json:"Пароль для сохранения"`
	NormalField5 NormalEntity95 `json:"Тип схемы"`
}

/*type IdEntity struct {
	Value string `json:"value"`
}

type TextEntity struct {
	Value string `json:"value"`
}*/

type NormalEntity90 struct {
	Value string `json:"value"`
}

type NormalEntity91 struct {
	Value string `json:"value"`
}

type NormalEntity92 struct {
	Value string `json:"value"`
}

type NormalEntity93 struct {
	Value string `json:"value"`
}

type NormalEntity94 struct {
	Value string `json:"value"`
}

type NormalEntity95 struct {
	Value string `json:"value"`
}

type ReferencesEntity struct {
	Class string `json:"class"`
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
