#!/usr/bin/env python3

import argparse
import csv
import subprocess
import os
from pathlib import Path

def run_kubectl_command(id_value):
    """
    Run kubectl exec command for the given ID and return the output.
    
    Args:
        id_value: The ID value from the CSV row
        
    Returns:
        The command output as a string
    """
    command = f"kubectl exec deployment/migration-backend -- like-migration-backend-cli likenft migrate-class-by-cosmos-class-id {id_value}  --premint-all-nfts"
    
    try:
        result = subprocess.run(
            command,
            shell=True,
            check=True,
            text=True,
            capture_output=True
        )
        return result.stdout
    except subprocess.CalledProcessError as e:
        print(f"Error executing command for ID {id_value}: {e}")
        return f"ERROR: {e.stderr}"

def process_csv(csv_file, id_column, output_dir=None):
    """
    Process each row in the CSV file, run kubectl command for each ID,
    and save output to individual files.
    
    Args:
        csv_file: Path to the CSV file
        id_column: Name of the ID column in the CSV
        output_dir: Directory to save output files (default: current directory)
    """
    if output_dir:
        os.makedirs(output_dir, exist_ok=True)
    
    with open(csv_file, 'r') as f:
        reader = csv.DictReader(f)
        
        if id_column not in reader.fieldnames:
            raise ValueError(f"Column '{id_column}' not found in CSV. Available columns: {reader.fieldnames}")
        
        for row in reader:
            id_value = row[id_column]
            print(f"Processing ID: {id_value}")
            
            output = run_kubectl_command(id_value)
            
            if output_dir:
                output_file = os.path.join(output_dir, f"{id_value}.out")
            else:
                output_file = f"{id_value}.out"
            
            with open(output_file, 'w') as out_f:
                out_f.write(output)
            
            print(f"Output saved to {output_file}")

def main():
    parser = argparse.ArgumentParser(description='Process CSV and run kubectl commands for each row')
    parser.add_argument('csv_file', help='Path to the CSV file')
    parser.add_argument('--id-column', default='ID', help='Name of the ID column in the CSV (default: ID)')
    parser.add_argument('--output-dir', help='Directory to save output files (default: current directory)')
    
    # Parse arguments
    args = parser.parse_args()
    
    # Process the CSV
    process_csv(args.csv_file, args.id_column, args.output_dir)

if __name__ == "__main__":
    main()
