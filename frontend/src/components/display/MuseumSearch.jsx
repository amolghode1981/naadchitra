import { useState } from "react";

export default function MuseumSearch({ museums, onFetch, loading }) {
  const [museum, setMuseum] = useState("met");
  const [id, setId] = useState("50939");

  function handleSubmit(e) {
    e.preventDefault();
    if (id.trim()) onFetch(museum, id.trim());
  }

  return (
    <form onSubmit={handleSubmit} className="flex gap-3 flex-wrap items-center">
      <label className="text-lapis font-fell italic text-sm sr-only">Museum</label>
      <select
        value={museum}
        onChange={(e) => setMuseum(e.target.value)}
        disabled={loading}
        className="bg-ivory border border-gold/50 text-mughal font-body text-sm
                   px-4 py-2.5 rounded focus:outline-none focus:border-gold
                   focus:ring-1 focus:ring-gold/30 disabled:opacity-50
                   shadow-sm uppercase tracking-widest cursor-pointer"
      >
        {museums.map((m) => (
          <option key={m} value={m}>{m.toUpperCase()}</option>
        ))}
      </select>

      <input
        type="text"
        value={id}
        onChange={(e) => setId(e.target.value)}
        placeholder="Object ID — e.g. 50939"
        disabled={loading}
        className="flex-1 min-w-52 bg-ivory border border-gold/50 text-mughal
                   font-body text-sm px-4 py-2.5 rounded focus:outline-none
                   focus:border-gold focus:ring-1 focus:ring-gold/30
                   placeholder:text-stone-400 disabled:opacity-50 shadow-sm"
      />

      <button
        type="submit"
        disabled={loading || !id.trim()}
        className="px-7 py-2.5 font-body text-sm font-medium tracking-widest uppercase
                   rounded transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed
                   shadow-sm"
        style={{
          background: "linear-gradient(135deg, #1a3a6e, #2a5298)",
          color: "#f0c040",
          border: "1px solid #c9a84c",
        }}
      >
        {loading ? "Searching…" : "Search Archive"}
      </button>
    </form>
  );
}
