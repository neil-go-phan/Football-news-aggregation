import { createContext } from "react";

interface IsLoggedContextType {
  isLogged: boolean;
    // eslint-disable-next-line no-unused-vars
  setIsLogged: (isLogged: boolean) => void;
}

export const IsLoggedContext = createContext<IsLoggedContextType>({
  isLogged: false,
  setIsLogged: () => {},
});
