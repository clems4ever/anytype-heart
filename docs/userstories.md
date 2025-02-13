### User stories

#### 0. Как ответы задавать
```js
// 1. Избыточная информация, плюс Result получится слишком кастомным
Package (id:'0x765', Reply {to:'0x123', error:Error{ type:WRONG_MNEMONIC, message:'Mnemonic is wrong' }, Result:{ type:FAILURE }})
// 2. Сообщение тоже получается ибыточным. из Error.type мы можем его получать на клиенте, плюс локализация
Package (id:'0x765', Status{ type:WRONG_MNEMONIC, message:'Mnemonic is wrong' })
// 3. 
Package (id:'0x765', Error{ type:WRONG_MNEMONIC })
Package (id:'0x765', Success{})
// 4. 
Package (id:'0x765', Status{ type:WRONG_MNEMONIC })
Package (id:'0x765', Status{ type:SUCCESS })
```

Выбрали четвертый вариант

#### 1. Log in

```js
// 1. Клиент передает мнемонику в middle, которую ввел пользователь
Front: Package (id:'0x123',  WalletLogin { mnemonic:'abc def ... xyz', pin:'12345'} )
Middle: Package (id:'0x980', Status { replyTo:'0x123', type:SUCCESS })
// 2. Middle начинает слать аккаунты
Middle: Package (id:'0x789', AccountFound { Account {name:'Pablo', id:'0xabcabc', icon:'0x123123'}}})
Middle: Package (id:'0x678', AccountFound{ Account {name:'Carlito', id:'0xabcabc', icon:'0x123123'}})
// 2.B. Middle сообщает об ошибке
Middle: Package (id:'0x765', Status { replyTo'0x123': type: WRONG_MNEMONIC })
// 3. Клиент отправляет аккаунт, под которым хочет работать
Front: Package (id:'0x789', AccountSelect {id:'0xabcabc'}})
Middle: Package (id:'0x777', Status { replyTo'0x789': type: SUCCESS })
```

#### 2. Sign up
```js
// 1. Просим создать аккаунт
Front: Package (id:'0x123', WalletCreate {} )
Middle: Package (id:'0x980', Status { replyTo'0x123': type: SUCCESS })

Front: Package (id:'0x345', AccountCreate { name:'Carlos', icon:'0x1231243257', pin:'1232724'} )
Middle: Package (id:'0x456', Status { replyTo'0x345': type: SUCCESS })
```

#### 3A. Получение списка документов (если store контролирует клиент)
Нужно получить список id документов, их имена, аватарки, хеши последних актуальных версий 
Когда нужен этот сценарий? Когда юзер хочет запустить главный экран.

1. Юзер запустил приложение. Middle уже авторизован, пока ничего не отрисовано
2. Фронт сообщает о том, какие у него документы есть 

```js
    Front: Message StartUp (docs: [
        {root:0x345, last_ver:0x123}, 
        {root:0x456, last_ver:0x234}, 
    ...])
```

3. Миддл сообщает, какие документы поменяли имена/аватарки, присылает их, актуальная ли версия хранимого документа, и если нет, то какая актуальная (или массив хешей CRDT-изменений, которые нужно скачать для восстановления до актуальной версии)

```js
    Middle: Message StartUp reply (docs: [
        {root:0x345, status:last_version}, 
        {root:0x456, status:outdated, name:same, icon:b64(newIcon.png), lastVersion:0x789},
    ...])
```

4. Клиент применяет полученные изменения и отображает список документов

#### 3B. Получение списка документов (если store контролирует middle)
Не вижу проблемы, если middle будет контролировать store. Плюсы – логика с клиента переходит на middle.

1. Юзер запустил приложение. Middle уже авторизован, пока ничего не отрисовано
2. Клиент сообщает, что он запустился

```js
    Front: Message StartUp ()
```

3. Middle отдает данные, которые нужно отрисовать на главной странице – список документов

```js
    Middle: Message DocumentsOrganizier (docs: [
        {name:'Doc 1', version:0x123, icon:icon1.png},
        {name:'Doc 2', version:0x234, icon:icon2.png},
    ...])
```

Логика по получению актуальных версий, сверки и прочего полностью абстрагирована от клиента.

4. Клиент просто отрисовывает полученные данные.

##### Cообщения сценария
1. Сообщение, которым клиент сообщает, что ему нужен отрисовать список документов. Возникает в сценариях, когда мы на главном меню, плюс, возможно, в других сценариях (например, какое-то всплывающее контекстное меню, в котором отображаются документы).
2. Сообщение, в котором middle передает список всех документов.

```js
    // С помощью запроса с entity == docHeaders можно запросить список документов
    // Выделять отдельное в сообщение DocumentsRequest не вижу смысла, оно слишком тривиальное получится
    message Request {
        string id = 0;
        string entity = 1;
        string target = 2;
    }

    // когда приходит DocHeaders, автоматом на фронте отрисовывается соответствующий target с docHeaders.
    message DocHeaders {
        string id = 0;
        repeated DocHeader docHeaders = 1;
    }

    message DocHeader {
        string id = 0;
        string name = 1;
        string root = 2;
        string version = 3;
        string iconName = 4;
    }
```

#### 4. Получение документа
1. Юзер находится в главном меню и видит список документов. Юзер нажимает на один из них
2. Клиент отправляет сообщение `Request { entity:document, target:0x123123 }`
3. Middle отправляет сообщение `Document { root:0x123123, ..., blocks:[...] }`
4. Клиент отрисовывает документ.