/** 空态引导 */
export function EmptyState({
  title,
  action,
}: {
  title: string;
  action?: React.ReactNode;
}) {
  return (
    <div className="glass-card flex flex-col items-center gap-4 px-6 py-14 text-center">
      <div className="flex h-12 w-12 items-center justify-center rounded-full bg-primary-soft text-2xl">
        ◈
      </div>
      <p className="text-text-secondary">{title}</p>
      {action}
    </div>
  );
}