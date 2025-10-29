use std::collections::HashMap;
use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader};

/*
Gaurav Sablok
codeprog@icloud.com
*/

#[derive(Debug, Clone, PartialEq, PartialOrd)]
pub struct FastaStruct {
    id: String,
    seq: String,
}

#[derive(Debug)]
pub struct FastaRecord {
    pub id: String,
    pub sequence: String,
}

#[tokio::main]
pub async fn read_transcript_fasta() -> Result<HashMap<String, FastaRecord>, Box<dyn Error>> {
    let pathfile = "./serverfiles/transcriptome.fasta";
    let file = File::open(pathfile)?;
    let reader = BufReader::new(file);
    let mut records = HashMap::new();
    let mut current_id = String::new();
    let mut current_sequence = String::new();
    for line in reader.lines() {
        let line = line?;
        if line.starts_with('>') {
            if !current_id.is_empty() {
                records.insert(
                    current_id.clone(),
                    FastaRecord {
                        id: current_id.clone(),
                        sequence: current_sequence.clone(),
                    },
                );
                current_sequence.clear();
            }
            current_id = line[1..].to_string();
        } else {
            current_sequence.push_str(&line);
        }
    }

    if !current_id.is_empty() {
        records.insert(
            current_id.clone(),
            FastaRecord {
                id: current_id,
                sequence: current_sequence,
            },
        );
    }

    Ok(records)
}
