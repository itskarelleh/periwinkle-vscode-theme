import { useState, useEffect } from "react";

const MOODS = ["grounded", "foggy", "tender", "light", "restless"];

function MoodSelector({ onSelect }) {
  const [active, setActive] = useState(null);

  const handleSelect = (mood) => {
    setActive(mood);
    onSelect(mood);
  };

  return (
    <div className="mood-selector">
      {MOODS.map((mood) => (
        <button
          key={mood}
          className={`mood-btn ${active === mood ? "active" : ""}`}
          onClick={() => handleSelect(mood)}
        >
          {mood}
        </button>
      ))}
    </div>
  );
}

function MirrorCard({ mood, insight }) {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => setVisible(true), 300);
    return () => clearTimeout(timer);
  }, [mood]);

  if (!mood) return null;

  return (
    <div className={`mirror-card ${visible ? "fade-in" : ""}`}>
      <h2 className="mirror-heading">your mirror</h2>
      <p className="mirror-mood">{mood}</p>
      {insight && <p className="mirror-insight">{insight}</p>}
    </div>
  );
}

export default function App() {
  const [currentMood, setCurrentMood] = useState(null);
  const [insight, setInsight] = useState("");

  const handleMoodSelect = async (mood) => {
    setCurrentMood(mood);
    // simulate fetch
    setInsight(`You chose ${mood}. That's worth sitting with.`);
  };

  return (
    <main className="app">
      <h1>zillinity</h1>
      <MoodSelector onSelect={handleMoodSelect} />
      <MirrorCard mood={currentMood} insight={insight} />
    </main>
  );
}