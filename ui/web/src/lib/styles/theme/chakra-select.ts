import { ChakraStylesConfig } from 'chakra-react-select'

export const chakraStylesConfig: ChakraStylesConfig = {
  menu: (provided) => ({
    ...provided,
    border: "1px solid",
  }),
}
