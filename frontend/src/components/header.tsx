import { useAuth } from "../context/auth";

export const Header = () => {
  const { user } = useAuth();

  return (
    <nav className="p-4 flex gap-4">
      <div className="flex gap-2 items-center">
        <div className="w-8 h-8 border border-gray-300 rounded-full bg-gray-100"></div>
        <span className="text-gray-900">{user?.name}</span>
      </div>
      <span>
        Today is {new Intl.DateTimeFormat("en-US").format(new Date())}
      </span>
    </nav>
  );
};
