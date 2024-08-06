# Archived as moved to yandex cloud object storage

## About [eng]
### Info
Here's the simplest platfrom to keep books: preview, author, title, download button - nothing more

Obviously I can't keep my pdfs on github (2 reasons :) ), so it's just a template 

The database is living in my library website - [図書館.きく.コム](図書館.きく.コム)


## Contribute [eng]
It would be great if you'd like to contribute a good book

### What to do?
Fork and add changes to the following files:
1) ```README.md``` - add your book's author and title in the following format:
```- [X] Author - Book Title```
2) ```templates/index.html``` - add div blocks as shown above in this file:
```html
<div class="book">
    <img src="/assets/somebooknameyougaveit.gif" alt="Error!">
    <details>
        <summary>Description</summary>
        <p>Author<br>Book Title<br><br><a
                href="/books/somebooknameyougaveit.pdf">Download</a></p>
    </details>
</div>
```
3) I'd be great if you'd also provide the link to the book :)
4) Pull Requst

Compress image before pushing ```jpegoptim --size=100k image.jpg -s```

Issues and better html code are also welcome :)

## About [ru]
Перед вами самая простая в мире платформа для хранения книжек: превью, автор, название, кнопка скачать - ничего лишнего

Очевидно я не могу хранить пдфки книг на гитхабе (по двум причинам :) ), поэтому это просто темплейт

Основная база живёт на моём сайте библиотеки - [図書館.きく.コム](図書館.きく.コム)

## Contribute [ru]
Будет замечательно, если вы предложите и поделитесь какой-то хорошей книжкой

### Что делать?
Сфоркать репозиторий себе и внести изменения в следующие файлы:
1) ```README.md``` - добавьте автора и название книги в следующем формате:
```- [X] Author - Book Title```
2) ```templates/index.html``` - добавьте блоки div, как показано в этом файле выше:
```html
<div class="book">
    <img src="/assets/somebooknameyougaveit.gif" alt="Error!">
    <details>
        <summary>Description</summary>
        <p>Author<br>Book Title<br><br><a
                href="/books/somebooknameyougaveit.pdf">Download</a></p>
    </details>
</div>
```
3) Будет также прекрасно, если вы поделитесь ссылкой на книгу :)
4) Пулл реквест

Картинку предварительно сжать ```jpegoptim --size=100k image.jpg -s```

Правки и более хороший html код очень приветствуются :)

## Docker
```bash
git clone https://github.com/TrueBad0ur/libreasy.git
docker compose up --build
access: 127.0.0.1:8080
```

## ToDo
- [ ] Add devops books
- [X] Add docker images / compose
- [ ] MB add pentest shelf

## Books

### Reverse Engineering

- [X] Enrico Martignetti - What Makes It Page?: The Windows 7 (x64) Virtual Memory Manager
- [X] Денис Юричев - Reverse Engineering для начинающих
- [X] Dmitry Vostokov - Memory Dump Analysis Anthology, Volume1
- [X] Dmitry Vostokov - Memory Dump Analysis Anthology, Volume2
- [X] Pavel Yosifovich - Windows Kernel Programming, Second Edition
- [X] Павел Йосифович - Работа с ядром Windows 1-ое издание
- [X] Руссинович М., Соломон Д., Ионеску А., Йосифович П - Внутреннее устройство Windows. 7-е изд
- [X] Руссинович М., Соломон Д. - Внутреннее устройство Microsoft Windows. 6-е изд, часть 2
- [X] Крис Касперски, Ева Рокко - Искусство дизассемблирования
- [X] Крис Касперски, Валентин Холмогоров, Ксения Кирилова - Восстановление данных. Практическое руководство
- [X] Mario Hewardt, Daniel Pravat - Advanced Windows Debugging 
- [X] Дармаван Салихан - BIOS - дизассемблеирование, модификация, программирование
- [X] Майкл Сикорски, Эндрю Хониг - Вскрытие покажет! Практический анализ вредоносного ПО
- [X] Massimiliano Tomassoli - Modern Windows Exploit Development
- [X] Свен Шрайбер - Недокументированные возможности Windows 2000
- [X] Алекс Матросов, Евгений Родионов, Сергей Братусь - Руткиты и буткиты. Обратная разработка вредоносных программ и угрозы следующего поколения
- [X] Хоглунд, Батлер - Руткиты: внедрение в ядро Windows
- [X] Mark Russinovich - Troubleshooting with the Windows Sysinternals Tools
- [X] Dennis Andriesse - Practical Binary Analysis: Build Your Own Linux Tools for Binary Instrumentation, Analysis, and Disassembly
- [X] Keith D. Cooper, Linda Torczon - Engineering: A Compiler
- [X] П.Л. Хижняк - Пишем вирус и антивирус  для совместимых компьютеров
- [X] Stephanie Domas, Christopher Domas - x86 Software Reverse-Engineering, Cracking, and Counter-Measures
- [X] Бруно Кардос Лопес, Рафаэль Аулер - LLVM: инфраструктура для разработки компиляторов
- [X] Ахо Альфред В., Ульман Джеффри Д., Лам Моника С., Сети Рави - Компиляторы: принципы, технологии и инструменты
- [X] Bill Blunden - The Rootkit Arsenal: Escape and Evasion: Escape and Evasion in the Dark Corners of the System
- [X] Бруно Кардос Лопес, Рафаэль Аулер - LLVM: инфраструктура для разработки компиляторов


### Cloud Engineering / DevOps

- [X] Автор Gerard Beekmans. Редакторы Matthew Burgess и Bruce Dubbs. Перевод Иван Лабутин - Linux From Scratch Версия 7.3


### Xakep
