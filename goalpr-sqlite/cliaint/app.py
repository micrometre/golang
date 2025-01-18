import sqlite3
import pandas as pd
connection = sqlite3.connect("../data/sqlite.db")

print(connection.total_changes)
cursor = connection.cursor()
rows = cursor.execute("SELECT plate, uuid FROM alprd").fetchall()
df = pd.DataFrame({'col':rows})
df.to_csv('../data/upload.csv', sep='\t', encoding='utf-8' )
(df.to_string()) 
print (df)
print (df)