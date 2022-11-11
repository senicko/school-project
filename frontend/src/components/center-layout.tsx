import { Header } from "./header";

export type CenterLayoutProps = {
  children: React.ReactNode;
};

export const CenterLayout = ({ children }: CenterLayoutProps) => {
  return (
    <main className="mx-auto flex max-w-5xl flex-col gap-16 p-4">
      <Header />
      {children}
    </main>
  );
};
