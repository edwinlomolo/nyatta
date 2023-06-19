import { Box, Flex, Text } from '@chakra-ui/react'

const Brand = ({ ...rest }: any): JSX.Element => (
  <Box
    {...rest}
    borderRight="1px"
    borderRightColor="gray.200"
    w={60}
    pos="fixed"
  >
    {/* TODO do mobile navigation */}
    <Flex alignItems="center" justifyContent="space-between" mx={8} h={20}>
      <Text fontSize="4xl" fontWeight="bold">Nyatta</Text>
    </Flex>
  </Box>
)

export default Brand
