// const Divider = () => <div className="h-[1px] w-full bg-gray-200" />;

export type ContextMenuProps = {
  onLogOut: () => void;
};

export const ContextMenu = ({ onLogOut }: ContextMenuProps) => {
  return (
    <div className="flex w-64 flex-col gap-4 rounded-lg border border-gray-200 bg-white p-4 shadow-sm">
      <button
        className="rounded-lg bg-red-500 p-2 text-white transition-all hover:bg-red-600"
        onClick={onLogOut}
      >
        Log Out
      </button>
    </div>
  );
};
