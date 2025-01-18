import sqlite3
import pandas as pd
import re

# Connect to the SQLite database
db_path = '../data/sqlite.db'  # Replace with the actual path
connection = sqlite3.connect(db_path)
print(connection.total_changes)

cursor = connection.cursor()
rows = cursor.execute("SELECT plate, uuid FROM alprd").fetchall()

# Create a pandas DataFrame
df = pd.DataFrame(rows, columns=['plate', 'uuid'])

# Define the regular expression for GB plates
test_gb = '^[A-Z]{2}[0-9]{2}[A-Z]{3}$'  # Corrected pattern

# Filter DataFrame for GB plates using boolean indexing
gb_plates = df[df['plate'].str.match(test_gb)]

# Create a new DataFrame with the filtered data
df2 = gb_plates[['plate', 'uuid']] 

# Save the new DataFrame to a CSV file
df2.to_csv('../data/gb-plates.csv', index=False)

print(df2)
print(df2)