import { LearningSetCreate } from "../components/learning-set/learning-set-create";
import { CenterLayout } from "../components/center-layout";
import { Header } from "../components/header";

export const AddWordSet = () => {
  return (
    <CenterLayout>
      <Header />
      <LearningSetCreate />
    </CenterLayout>
  );
};
