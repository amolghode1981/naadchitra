import { useState, useEffect } from "react";
import MuseumSearch from "./components/display/MuseumSearch";
import ArtworkCard from "./components/display/ArtworkCard";
import ErrorMessage from "./components/display/ErrorMessage";

const API_BASE = "http://localhost:8080";

export default function App() {
  const [museums, setMuseums] = useState([]);
  const [artwork, setArtwork] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(`${API_BASE}/api/museums`)
      .then((r) => r.json())
      .then(setMuseums)
      .catch(() => setMuseums(["met"]));
  }, []);

  async function handleFetch(museum, id) {
    setLoading(true);
    setError(null);
    setArtwork(null);
    try {
      const res = await fetch(
        `${API_BASE}/api/object?museum=${encodeURIComponent(museum)}&id=${encodeURIComponent(id)}`
      );
      if (!res.ok) throw new Error(await res.text());
      setArtwork(await res.json());
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="min-h-screen bg-ivory bg-miniature-pattern font-body">

      {/* Top decorative border */}
      <div className="h-1.5 w-full" style={{background: "linear-gradient(90deg, #1a3a6e, #c9a84c, #c8570a, #c9a84c, #1a3a6e)"}} />

      {/* Header */}
      <header className="bg-ivory-dark border-b border-gold/40 px-6 py-8">
        <div className="max-w-5xl mx-auto text-center">
          {/* Ornamental top line */}
          <div className="flex items-center justify-center gap-3 mb-4">
            <div className="h-px flex-1 max-w-24" style={{background: "linear-gradient(90deg, transparent, #c9a84c)"}} />
            <span className="text-gold text-lg">✦</span>
            <div className="h-px flex-1 max-w-24" style={{background: "linear-gradient(90deg, #c9a84c, transparent)"}} />
          </div>

          <p className="font-fell italic text-saffron text-sm tracking-widest uppercase mb-2">
            Collection Archive
          </p>
          <h1 className="font-display text-5xl font-semibold text-mughal tracking-tight leading-tight">
            Museum of Indian Miniatures
          </h1>
          <p className="font-fell italic text-lapis mt-2 text-lg">
            Paintings from the courts of the Mughal, Rajput & Pahari traditions
          </p>

          {/* Ornamental bottom line */}
          <div className="flex items-center justify-center gap-3 mt-4">
            <div className="h-px flex-1 max-w-24" style={{background: "linear-gradient(90deg, transparent, #c9a84c)"}} />
            <span className="text-gold text-lg">✦</span>
            <div className="h-px flex-1 max-w-24" style={{background: "linear-gradient(90deg, #c9a84c, transparent)"}} />
          </div>
        </div>
      </header>

      {/* Search */}
      <div className="bg-lapis/5 border-b border-gold/20 px-6 py-6">
        <div className="max-w-5xl mx-auto">
          <MuseumSearch museums={museums} onFetch={handleFetch} loading={loading} />
        </div>
      </div>

      {/* Main content */}
      <main className="max-w-5xl mx-auto px-6 py-10">
        {error && <ErrorMessage message={error} />}

        {loading && (
          <div className="flex flex-col items-center justify-center py-24 gap-4">
            <div className="w-10 h-10 rounded-full border-2 border-gold/30 border-t-gold animate-spin" />
            <p className="font-fell italic text-saffron text-sm tracking-widest">
              Retrieving from the archive…
            </p>
          </div>
        )}

        {artwork && <ArtworkCard artwork={artwork} />}

        {!artwork && !loading && !error && (
          <div className="text-center py-24">
            <p className="font-fell italic text-lapis/50 text-xl">
              Enter an object ID above to begin exploring the collection
            </p>
            <p className="text-stone-400 text-sm mt-2 font-body">
              Try object ID <span className="font-mono bg-gold/10 px-2 py-0.5 rounded text-saffron">50939</span> to start
            </p>
          </div>
        )}
      </main>

      {/* Bottom border */}
      <div className="h-1.5 w-full mt-auto" style={{background: "linear-gradient(90deg, #1a3a6e, #c9a84c, #c8570a, #c9a84c, #1a3a6e)"}} />
    </div>
  );
}
