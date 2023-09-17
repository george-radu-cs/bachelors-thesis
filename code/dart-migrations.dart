const Map<int, String> migrationsScripts = {
  1: '''
    CREATE TABLE settings (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      languageCode TEXT NOT NULL,
      themeMode TEXT NOT NULL,
      themeColor TEXT NOT NULL
    );
  ''',
  2: '''
    INSERT INTO settings (languageCode, themeMode, themeColor) 
      VALUES ('en', 'system', 'blue');
  ''',
  // ...
  11: '''
    ALTER TABLE settings ADD COLUMN showOCRDialog BOOLEAN NOT NULL DEFAULT 1;
  ''',
}