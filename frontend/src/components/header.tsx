import ProfileImage from "../assets/profile.png";
import { useUserStore } from "../state/user";

export const Header = () => {
  const user = useUserStore((state) => state.user);

  return (
    <nav className="flex items-center gap-4 justify-between">
      <div className="flex gap-4 items-center">
        <img
          src={ProfileImage}
          alt="profile image"
          className="rounded-full aspect-square w-[48px] border border-gray-800"
        />
        <span className="text-gray-800">{user?.name}</span>
      </div>
    </nav>
  );
};
