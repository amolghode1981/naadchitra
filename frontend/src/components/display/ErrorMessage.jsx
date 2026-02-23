export default function ErrorMessage({ message }) {
  return (
    <div
      className="flex items-start gap-3 px-5 py-4 rounded mb-8 font-body text-sm"
      style={{
        background: "rgba(192,57,43,0.06)",
        border: "1px solid rgba(192,57,43,0.3)",
        color: "#922b21",
      }}
    >
      <span className="text-vermil font-bold mt-0.5 shrink-0">✕</span>
      <p>{message}</p>
    </div>
  );
}
