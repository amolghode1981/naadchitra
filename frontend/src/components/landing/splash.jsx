export default function Landing({ onStart }) {
  function handleClick() {
    onStart();
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-4xl font-display text-mughal mb-4">
        Naad Chitra : Imaging the sound
      </h1>
      <p className="text-stone-500 font-body mb-8">
        Mapping Musical Activity in Indian Miniature Painting
      </p>
      <button
        onClick={handleClick}
        className="px-8 py-3 font-body text-sm font-medium tracking-widest
                   uppercase rounded transition-all duration-200"
        style={{
          background: "linear-gradient(135deg, #1a3a6e, #2a5298)",
          color: "#f0c040",
          border: "1px solid #c9a84c",
        }}
      >
        Explore Collection
      </button>
    </div>
  );
}
