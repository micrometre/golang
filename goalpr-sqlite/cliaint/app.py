import sqlite3
import pandas as pd
import re

# Connect to the SQLite database
connection = sqlite3.connect('../data/sqlite.db')
print(connection.total_changes)

cursor = connection.cursor()
rows = cursor.execute("SELECT plate, uuid FROM alprd").fetchall()

# Create a pandas DataFrame
df = pd.DataFrame(rows, columns=['plate', 'uuid'])

# Define the regular expression for GB plates
test_gb = '^[A-Z]{2}[0-9]{2}[A-Z]{3}$'  # Corrected pattern

# Filter DataFrame for GB plates using boolean indexing
gb_plates = df[df['plate'].str.match(test_gb)]

# Create a dictionary with plate as key and uuid as value from the filtered DataFrame
plates = gb_plates.set_index('plate')['uuid'].to_dict()

print(plates)
#df.to_csv('../data/upload.csv', sep='\t', encoding='utf-8' )