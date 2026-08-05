import { COPY } from "@/lib/copy";
import { USE_MOCK } from "@/lib/config";

/** 页脚:固定合规披露(文档 §5.3) */
export function Footer() {
  return (
    <footer className="mt-12 border-t border-border/60 py-6">
      <div className="mx-auto flex max-w-3xl flex-col items-center gap-2 px-4 text-center">
        <ul className="flex flex-wrap items-center justify-center gap-x-4 gap-y-1 text-xs text-text-muted">
          {COPY.footer.map((line) => (
            <li key={line}>{line}</li>
          ))}
        </ul>
        {USE_MOCK && (
          <span className="rounded-full border border-shortfall/40 bg-shortfall/10 px-2.5 py-0.5 text-[11px] text-shortfall">
            {COPY.demoBadge}:当前界面使用演示数据,对接后端后自动关闭
          </span>
        )}
      </div>
    </footer>
  );
}