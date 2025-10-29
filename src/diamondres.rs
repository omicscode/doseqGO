use std::collections::HashMap;
use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader};

/*
Gaurav Sablok
codeprog@icloud.com
*/

#[derive(Debug, Clone, PartialOrd, PartialEq)]
pub struct GenomeStruct {
    pub id: String,
    pub annotation: String,
    pub identity: String,
}

#[tokio::main]
pub async fn diamonresult() -> Result<String, Box<dyn Error>> {
    let pathfile = "./serverfiles/diamondresult.txt";
    let mut idhashmap: Vec<_> = HashMap::new();
    let mut idmap: Vec<(String, String)> = Vec::new();
    let fileopen = File::open(pathfile).expect("File not found");
    let fileread = BufReader::new(pathfile);

    for i in fileread.lines() {
        let line = i.expect("line not present");
        let linevec = line.split("\t").map(|x| x.to_string()).collect::<Vec<_>>();
        idhashmap.push(linevec[0].clone());
        idmap.push((linevec[0].clone, linevec[1..].concat()));
    }

    let mut vecnew: Vec<GenomeStruct> = Vec::new();

    for i in idhashmap.iter() {
        for j in idmap.iter() {
            if i == j.0 {
                vecnew.push(GenomeStruct {
                    id: i.clone,
                    annotation: j.1.split(" ").collect::<Vec<_>>()[0],
                    identity: j.1.split(" ").collect::<Vec<_>>()[1],
                });
            }
        }
    }

    Ok("The function has completed".to_string())
}
