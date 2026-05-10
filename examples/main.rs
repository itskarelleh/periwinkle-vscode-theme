use std::collections::HashMap;

#[derive(Debug, Clone)]
pub enum Mood {
    Grounded,
    Foggy,
    Tender,
    Light,
    Restless,
    Heavy,
    Still,
    Raw,
    Calm,
}

impl Mood {
    pub fn label(&self) -> &'static str {
        match self {
            Mood::Grounded => "grounded",
            Mood::Foggy => "foggy",
            Mood::Tender => "tender",
            Mood::Light => "light",
            Mood::Restless => "restless",
            Mood::Heavy => "heavy",
            Mood::Still => "still",
            Mood::Raw => "raw",
            Mood::Calm => "calm",
        }
    }

    pub fn intensity(&self) -> u8 {
        match self {
            Mood::Heavy | Mood::Raw => 9,
            Mood::Restless | Mood::Foggy => 6,
            Mood::Tender | Mood::Grounded => 4,
            Mood::Light | Mood::Calm | Mood::Still => 2,
        }
    }
}

#[derive(Debug)]
pub struct JournalEntry {
    pub mood: Mood,
    pub body: String,
    pub wins: Vec<String>,
    pub timestamp: u64,
}

impl JournalEntry {
    pub fn new(mood: Mood, body: impl Into<String>) -> Self {
        Self {
            mood,
            body: body.into(),
            wins: Vec::new(),
            timestamp: 0, // replace with real unix ts
        }
    }

    pub fn add_win(&mut self, win: impl Into<String>) {
        self.wins.push(win.into());
    }

    pub fn summary(&self) -> String {
        format!(
            "[{}] {} — {} wins logged",
            self.mood.label(),
            &self.body[..self.body.len().min(40)],
            self.wins.len()
        )
    }
}

fn seed_from_mood(mood: &Mood) -> u64 {
    // FNV-1a inspired deterministic seed
    let label = mood.label();
    let mut hash: u64 = 0xcbf29ce484222325;
    for byte in label.bytes() {
        hash ^= byte as u64;
        hash = hash.wrapping_mul(0x100000001b3);
    }
    hash
}

fn main() {
    let mut entry = JournalEntry::new(Mood::Tender, "sat with the feeling instead of running");
    entry.add_win("shipped The Record feature");
    entry.add_win("woke up at 4am and actually started");

    println!("{}", entry.summary());
    println!("seed: {}", seed_from_mood(&entry.mood));

    let moods = vec![Mood::Calm, Mood::Restless, Mood::Raw];
    let intensities: HashMap<&str, u8> = moods
        .iter()
        .map(|m| (m.label(), m.intensity()))
        .collect();

    for (label, intensity) in &intensities {
        println!("{}: intensity {}", label, intensity);
    }
}