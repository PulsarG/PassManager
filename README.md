Password manager with asynchronous encryption/decryption.

Unlike other password managers, EnigmaPass keeps the database encrypted at all times, and decrypts only one position at a time (no more), and only in two cases:

- when copying login or password to clipboard

- when showing the login and password of one of the positions (no more than one at a time).

Moreover, thanks to asynchronous encryption, positions with logins and passwords are not tied to a single master password from the database.

This means that by opening the database, you can create new positions / groups, each with its own password, without restarting the application and the database.
And also in the reverse order to decrypt the positions / groups each with their own personal password in real time.
And even better: use different master passwords for login and password. It all depends on your cunning.

As long as one master password is valid, positions to which it does not apply will not be decrypted, even if they were previously decrypted by their own master password.

The master password is valid for the entire database, but CORRECTLY decrypts only those positions to which it was assigned by the user.

+

Additional protection is the generation of a unique "Rotor".

Basically, it's a key file. But it does not just help to decrypt the base, but contains the correct rules, "paths" to decryption. 

Without it, you will not be able to get the correct data even with the correct password. 

Each click on "Generate Custom Rotor" will create a unique file for each user.

# РУС

Менеджер паролей с асинхронным шифрованием/дешифрованием.

В отличие от других менеджеров паролей, EnigmaPass постоянно держит базу данных зашифрованной и расшифровывает только одну позицию за раз (не более), и только в двух случаях:

при копировании логина или пароля в буфер обмена;

при отображении логина и пароля одной из позиций (не более одной за раз).

Более того, благодаря асинхронному шифрованию, позиции с логинами и паролями не привязаны к одному единственному мастер-паролю из базы данных.

Это означает, что, открыв базу данных, вы можете создавать новые позиции/группы, каждая со своим собственным паролем, без перезапуска приложения и базы данных. А также в обратном порядке, чтобы расшифровывать позиции/группы, каждая со своим личным паролем, в режиме реального времени. И еще лучше: используйте разные мастер-пароли для логина и пароля. Все зависит от вашей хитрости.

Пока действителен хотя бы один мастер-пароль, позиции, к которым он не применяется, не будут расшифрованы, даже если они ранее были расшифрованы своим собственным мастер-паролем.

Главный пароль действителен для всей базы данных, но ПРАВИЛЬНО расшифровывает только те позиции, которые были ему назначены пользователем.

Дополнительной защитой является генерация уникального «Ротора».

По сути, это файл ключа. Но он не просто помогает расшифровать базу данных, а содержит правильные правила, «пути» к расшифровке.

Без него вы не сможете получить правильные данные, даже имея правильный пароль.

Каждый клик по кнопке «Создать пользовательский Ротор» создаст уникальный файл для каждого пользователя.


