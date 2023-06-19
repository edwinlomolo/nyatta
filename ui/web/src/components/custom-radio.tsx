import { Icon } from '@chakra-ui/icons'
import { Box, VStack, HStack, UseRadioProps, useRadio, Text } from '@chakra-ui/react'
import { FaCheckCircle, FaRegCircle } from 'react-icons/fa'

const CustomRadio = (props: UseRadioProps & { label: string, description?: string }) => {
  const { state, getInputProps, getCheckboxProps } = useRadio(props)
  const input = getInputProps()
  const checkbox = getCheckboxProps()
  const isChecked = state.isChecked

  return (
    <Box as="label" mb={5}>
      <input {...input} />
      <VStack
        {...checkbox}
        cursor="pointer"
        border="1px solid"
        borderRadius="md"
        _checked={{
          color: "green"
        }}
        shadow={isChecked ? 'md' : undefined}
        align="start"
        p={4}
        _disabled={{
          opacity: 0.4,
          cursor: 'not-allowed',
        }}
      >
        <HStack spacing={4} justify="space-between" w="full">
          <Text fontWeight="bold">{props.label}</Text>
          {isChecked ? <Icon as={FaCheckCircle} /> : <Icon as={FaRegCircle} />}
        </HStack>
        {props.description && <Text fontSize="sm">{props.description}</Text>}
      </VStack>
    </Box>
  )
}

export default CustomRadio
