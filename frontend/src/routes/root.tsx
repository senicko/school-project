import { CenterLayout } from "../components/center-layout";
import { JokesGenerator } from "../components/jokes/jokes-generator";
import { SavedJokes } from "../components/jokes/saved-jokes";

export const Root = () => {
  return (
    <CenterLayout>
      <JokesGenerator />
      <SavedJokes />
    </CenterLayout>
  );
};
