use std::collections::{HashMap, HashSet};
use std::error::Error;
use std::fmt;
use std::io::{BufRead, BufReader};
use std::str::FromStr;

/*
Gaurav Sablok
codeprog@icloud.com
*/

#[derive(Debug)]
pub struct ParseError {
    message: String,
}

impl fmt::Display for ParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Parse error: {}", self.message)
    }
}

impl Error for ParseError {}

impl ParseError {
    fn new(msg: &str) -> Self {
        ParseError {
            message: msg.to_string(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct Node {
    pub id: String,
    pub sequence: String,
    pub species: String,
    pub start_pos: i32,
    pub strand_rank: i32,
}

#[derive(Debug)]
pub struct Graph {
    pub nodes: HashMap<String, Node>,
    pub edges: HashMap<String, HashSet<String>>,
}

impl Graph {
    pub fn new() -> Self {
        Graph {
            nodes: HashMap::new(),
            edges: HashMap::new(),
        }
    }

    pub fn add_node(&mut self, id: String, sequence: String, species: String, start: i32, sr: i32) {
        self.nodes.insert(
            id.clone(),
            Node {
                id,
                sequence,
                species,
                start_pos: start,
                strand_rank: sr,
            },
        );
    }

    pub fn add_edge(&mut self, from: String, to: String) {
        self.edges
            .entry(from)
            .or_insert_with(HashSet::new)
            .insert(to);
    }

    pub fn parse<R: BufRead>(reader: R) -> Result<Self, Box<dyn Error>> {
        let mut graph = Graph::new();
        let scanner = BufReader::new(reader);

        for line_result in scanner.lines() {
            let line = line_result?;
            let line = line.trim();

            if line.is_empty() || line.starts_with('#') {
                continue;
            }

            let fields: Vec<&str> = line.split_whitespace().collect();
            if fields.is_empty() {
                continue;
            }

            let record_type = fields[0];

            match record_type {
                "S" => {
                    if fields.len() < 5 {
                        return Err(Box::new(ParseError::new(&format!(
                            "invalid S line: {}",
                            line
                        ))));
                    }

                    let id = fields[1].to_string();
                    let seq = fields[2].to_string();
                    let mut species = String::new();
                    let mut start_pos = 0;
                    let mut sr = 0;

                    for i in 3..fields.len() {
                        let parts: Vec<&str> = fields[i].splitn(3, ':').collect();
                        if parts.len() != 3 {
                            continue;
                        }
                        let tag = parts[0];
                        let typ = parts[1];
                        let val = parts[2];

                        match tag {
                            "SN" if typ == "Z" => species = val.to_string(),
                            "SO" if typ == "i" => {
                                if let Ok(v) = i32::from_str(val) {
                                    start_pos = v;
                                }
                            }
                            "SR" if typ == "i" => {
                                if let Ok(v) = i32::from_str(val) {
                                    sr = v;
                                }
                            }
                            _ => {}
                        }
                    }

                    graph.add_node(id, seq, species, start_pos, sr);
                }
                "L" => {
                    if fields.len() < 6 {
                        return Err(Box::new(ParseError::new(&format!(
                            "invalid L line: {}",
                            line
                        ))));
                    }

                    let from = fields[1];
                    let from_orient = fields[2];
                    let to = fields[3];
                    let to_orient = fields[4];

                    let edge_from = format!("{}{}", from, from_orient);
                    let edge_to = format!("{}{}", to, to_orient);

                    graph.add_edge(edge_from, edge_to);
                }
                _ => {}
            }
        }

        Ok(graph)
    }

    pub fn print_graph(&self) {
        println!("=== NODES ===");
        for (id, node) in &self.nodes {
            let orient = if node.strand_rank == 1 { "-" } else { "+" };
            println!(
                "{} ({}) [{}] pos:{} len:{}",
                id,
                orient,
                node.species,
                node.start_pos,
                node.sequence.len()
            );
        }

        println!("\n=== EDGES ===");
        for (from, targets) in &self.edges {
            for to in targets {
                println!("{} -> {}", from, to);
            }
        }
    }

    pub fn dfs(&self, start: &str) -> Vec<Vec<String>> {
        let mut results = Vec::new();
        let mut visited = HashMap::new();
        let mut path = Vec::new();

        self.dfs_recursive(start, &mut visited, &mut path, &mut results);
        results
    }

    fn dfs_recursive<'a>(
        &'a self,
        current: &'a str,
        visited: &mut HashMap<&'a str, bool>,
        path: &mut Vec<String>,
        results: &mut Vec<Vec<String>>,
    ) {
        visited.insert(current, true);
        path.push(current.to_string());

        let neighbors = self.edges.get(current).unwrap_or(&HashSet::new().clone());

        if neighbors.is_empty() {
            results.push(path.clone());
        } else {
            for neighbor in neighbors {
                if !visited.get(neighbor.as_str()).unwrap_or(&false) {
                    self.dfs_recursive(neighbor, visited, path, results);
                }
            }
        }
        path.pop();
        visited.insert(current, false);
    }
}

// Unit tests
#[cfg(test)]
mod tests {
    use super::*;
    use std::io::Cursor;

    #[test]
    fn test_parse_graph() {
        let input = r#"
        S	1	ACGT	SN:Z:Homo_sapiens	SO:i:100	SR:i:0
        S	2	TGCA	SN:Z:Homo_sapiens	SO:i:200	SR:i:1
        L	1	+	2	-
        "#;

        let cursor = Cursor::new(input.as_bytes());
        let graph = Graph::parse(cursor).unwrap();

        assert_eq!(graph.nodes.len(), 2);
        assert!(graph.nodes.contains_key("1"));
        assert!(graph.nodes.contains_key("2"));
        assert_eq!(graph.nodes["1"].sequence, "ACGT");
        assert_eq!(graph.nodes["2"].strand_rank, 1);

        let edges = graph.edges.get("1+").unwrap();
        assert!(edges.contains("2-"));
    }

    #[test]
    fn test_dfs() {
        let mut graph = Graph::new();
        graph.add_node("1".to_string(), "A".to_string(), "Human".to_string(), 0, 0);
        graph.add_node("2".to_string(), "B".to_string(), "Human".to_string(), 0, 0);
        graph.add_node("3".to_string(), "C".to_string(), "Human".to_string(), 0, 0);

        graph.add_edge("1+".to_string(), "2+".to_string());
        graph.add_edge("2+".to_string(), "3+".to_string());

        let paths = graph.dfs("1+");
        assert_eq!(paths.len(), 1);
        assert_eq!(paths[0], vec!["1+", "2+", "3+"]);
    }
}
