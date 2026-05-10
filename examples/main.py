from dataclasses import dataclass, field
from typing import Optional
from enum import Enum
import hashlib
import time


class JournalMode(Enum):
    FLOAT = "float"
    SWIM = "swim"


MOODS = [
    "grounded", "foggy", "tender", "light",
    "restless", "heavy", "still", "raw", "calm"
]


@dataclass
class Win:
    title: str
    timestamp: float = field(default_factory=time.time)

    def __str__(self) -> str:
        return f"+ {self.title}"


@dataclass
class JournalEntry:
    mood: str
    body: str
    mode: JournalMode = JournalMode.FLOAT
    wins: list[Win] = field(default_factory=list)
    timestamp: float = field(default_factory=time.time)

    def __post_init__(self):
        if self.mood not in MOODS:
            raise ValueError(f"Invalid mood: {self.mood}. Must be one of {MOODS}")

    def add_win(self, title: str) -> None:
        self.wins.append(Win(title=title))

    def summary(self) -> str:
        preview = self.body[:60] + "..." if len(self.body) > 60 else self.body
        return f"[{self.mood}] {preview} ({len(self.wins)} wins)"


def generate_blob_seed(mood: str, user_id: str) -> int:
    """Generate a deterministic seed for Mirror blob visuals."""
    raw = f"{mood}:{user_id}"
    digest = hashlib.md5(raw.encode()).hexdigest()
    return int(digest[:8], 16)


def get_insight(entry: JournalEntry, model: Optional[str] = None) -> str:
    """Stub for Mirror AI insight generation."""
    if not entry.body:
        return "Nothing to reflect on yet."
    word_count = len(entry.body.split())
    return (
        f"You wrote {word_count} words in {entry.mode.value} mode "
        f"while feeling {entry.mood}."
    )


if __name__ == "__main__":
    entry = JournalEntry(
        mood="tender",
        body="the morning was quiet enough to hear myself think",
        mode=JournalMode.SWIM,
    )
    entry.add_win("finished testing The Record")
    entry.add_win("published Substack piece")

    print(entry.summary())
    print(get_insight(entry))

    seed = generate_blob_seed(entry.mood, "user_karelle")
    print(f"blob seed: {seed}")

    for win in entry.wins:
        print(win)