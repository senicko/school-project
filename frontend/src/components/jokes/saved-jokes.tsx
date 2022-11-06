import { useUserStore } from "../../state/user";

export const SavedJokes = () => {
  const user = useUserStore((state) => state.user);

  return (
    <div>
      <h2 className="font-medium text-lg">Saved jokes</h2>
      {user && user.jokes.map((joke) => <p key={joke}>{joke}</p>)}
    </div>
  );
};
