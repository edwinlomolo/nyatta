import { FormControl, SimpleGrid, FormErrorMessage, FormLabel, useRadioGroup } from '@chakra-ui/react'
import { useController } from 'react-hook-form'

import CustomRadio from './custom-radio'

const FormRadio = ({ control, columns=2, options, name, label }: any): JSX.Element => {
  const { field, fieldState: { error } } = useController({
    name,
    control,
  })
  const { getRootProps, getRadioProps } = useRadioGroup(field)

  return (
    <FormControl isInvalid={!!error}>
      {label && <FormLabel htmlFor={name}>{label}</FormLabel>}
      <SimpleGrid {...getRootProps()} columns={columns} spacing={4}>
        {options.map(({ label, value, description }: any) => (
          <CustomRadio
            label={label}
            description={description}
            key={value}
            {...getRadioProps({ value })}
          />
        ))}
      </SimpleGrid>
      {error && <FormErrorMessage mb={5}>{error.message}</FormErrorMessage>}
    </FormControl>
  )
}

export default FormRadio
