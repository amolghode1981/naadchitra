export default function ArtworkCard({ artwork }) {
  const fields = [
    { label: "Date",       value: artwork.date,       icon: "◈" },
    { label: "Medium",     value: artwork.medium,     icon: "◈" },
    { label: "Dimensions", value: artwork.dimensions, icon: "◈" },
    { label: "Department", value: artwork.department, icon: "◈" },
    { label: "Culture",    value: artwork.culture,    icon: "◈" },
    { label: "Museum",     value: artwork.museum,     icon: "◈" },
  ].filter((f) => f.value);

  return (
    <div className="animate-in fade-in duration-700">

      {/* Split layout: image left, details right */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-10 items-start">

        {/* Left: Image */}
        {artwork.imageUrl && (
          <div
            className="rounded overflow-hidden shadow-mughal"
            style={{ border: "2px solid #c9a84c" }}
          >
            {/* Inner ornamental border */}
            <div className="p-2" style={{background: "linear-gradient(135deg, #fdf6e3, #f5e6c8)"}}>
              <div className="p-1" style={{border: "1px solid rgba(201,168,76,0.4)"}}>
                <img
                  src={artwork.imageUrl}
                  alt={artwork.title}
                  className="w-full object-contain max-h-[520px]"
                />
              </div>
            </div>
          </div>
        )}

        {/* Right: Details */}
        <div>
          {/* Title block */}
          <div className="mb-6 pb-6" style={{borderBottom: "1px solid rgba(201,168,76,0.3)"}}>
            <p className="font-fell italic text-saffron text-xs tracking-widest uppercase mb-2">
              {artwork.department || "Collection Object"}
            </p>
            <h2 className="font-display text-4xl font-semibold text-mughal leading-tight mb-3">
              {artwork.title || "Untitled"}
            </h2>
            {artwork.artist && (
              <p className="font-fell italic text-lapis text-xl">
                {artwork.artist}
              </p>
            )}
          </div>

          {/* Metadata fields */}
          <div className="space-y-4 mb-8">
            {fields.map(({ label, value }) => (
              <div key={label} className="flex gap-4">
                <span className="text-gold text-xs mt-1 shrink-0">◈</span>
                <div>
                  <p className="text-saffron/70 text-xs uppercase tracking-widest font-body font-medium mb-0.5">
                    {label}
                  </p>
                  <p className="text-mughal text-sm leading-relaxed font-body">
                    {value}
                  </p>
                </div>
              </div>
            ))}
          </div>

          {/* Tags */}
          {artwork.tags?.length > 0 && (
            <div className="mb-8">
              <p className="text-saffron/70 text-xs uppercase tracking-widest font-body font-medium mb-3">
                Subjects & Themes
              </p>
              <div className="flex flex-wrap gap-2">
                {artwork.tags.map((tag) => (
                  <span
                    key={tag}
                    className="text-xs font-body px-3 py-1 rounded-full text-lapis"
                    style={{
                      background: "rgba(26,58,110,0.08)",
                      border: "1px solid rgba(26,58,110,0.2)",
                    }}
                  >
                    {tag}
                  </span>
                ))}
              </div>
            </div>
          )}

          {/* Source link */}
          {artwork.sourceUrl && (
            <a
              href={artwork.sourceUrl}
              target="_blank"
              rel="noreferrer"
              className="inline-flex items-center gap-2 font-fell italic text-sm
                         text-saffron hover:text-saffron-dark transition-colors
                         underline underline-offset-4 decoration-gold/40"
            >
              View in museum archive →
            </a>
          )}
        </div>
      </div>

      {/* Bottom ornament */}
      <div className="flex items-center justify-center gap-3 mt-12">
        <div className="h-px flex-1" style={{background: "linear-gradient(90deg, transparent, #c9a84c)"}} />
        <span className="text-gold text-sm">✦</span>
        <span className="font-fell italic text-stone-400 text-xs tracking-widest">
          {artwork.museum}
        </span>
        <span className="text-gold text-sm">✦</span>
        <div className="h-px flex-1" style={{background: "linear-gradient(90deg, #c9a84c, transparent)"}} />
      </div>
    </div>
  );
}
