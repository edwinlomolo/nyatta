import { Button, FormControl, Flex, FormHelperText, IconButton, Input, VStack, Modal, ModalBody, ModalCloseButton, ModalHeader, ModalContent, useDisclosure } from '@chakra-ui/react'
import { Select } from 'chakra-react-select'
import { Controller, useForm } from 'react-hook-form'
import { BsFilterSquareFill } from 'react-icons/bs'

import data from 'data/data.json'
import { SearchListingsForm } from 'types'

const SearchForm = (): JSX.Element => {
	const { control, register, reset } = useForm<SearchListingsForm>()
	const { isOpen, onOpen, onClose } = useDisclosure()

	return (
		<Flex>
			<IconButton
				aria-label="Search database"
				icon={<BsFilterSquareFill />}
				variant="ghost"
				colorScheme="teal"
				size="lg"
				onClick={onOpen}
			/>
			<Modal motionPreset="slideInBottom" isOpen={isOpen} onClose={onClose}>
				<ModalContent>
					<ModalHeader>Search Listings</ModalHeader>
					<ModalCloseButton />
					<ModalBody>
						<VStack spacing={{base: 4, md: 6}}>
							<FormControl>
								<Input {...register('town')} placeholder="Town" />
								<FormHelperText>Town</FormHelperText>
							</FormControl>

							<FormControl>
								<Input {...register('minPrice')} placeholder="Minimum rent" />
								<FormHelperText>Minimum rent</FormHelperText>
							</FormControl>
							
							<FormControl>
								<Input {...register('maxPrice')} placeholder="Maximum rent" />
								<FormHelperText>Maximum rent</FormHelperText>
							</FormControl>

							<FormControl>
								<Controller
									name="propertyType"
									control={control}
									render={({ field }) => (
										<Select
											{...field}
											placeholder="Type..."
											isMulti
											isClearable
											options={data.propertyTypes as any[]}
										/>
									)}
								/>
								<FormHelperText>Select property type</FormHelperText>
							</FormControl>
							<Flex w="100%" justifyContent="space-between">
								<Button onClick={() => reset()}>Clear</Button>
								<Button>Found(0)</Button>
							</Flex>
						</VStack>
					</ModalBody>
				</ModalContent>
			</Modal>
		</Flex>
	)
}

export default SearchForm
