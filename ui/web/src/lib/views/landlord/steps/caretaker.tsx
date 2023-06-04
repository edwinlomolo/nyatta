import React, { type Dispatch, type SetStateAction } from 'react'

import { useMutation } from '@apollo/client'

import { Box, Center, Button, HStack, Image, FormControl, FormErrorMessage, FormHelperText, FormLabel, Icon, Input, Modal, ModalCloseButton, ModalHeader, ModalContent, ModalBody, Spacer, Stack, Textarea, useDisclosure, Spinner } from '@chakra-ui/react'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { yupResolver } from '@hookform/resolvers/yup'
import { useDropzone } from 'react-dropzone'
import { FaUpload } from 'react-icons/fa'

import { useForm, type SubmitHandler } from 'react-hook-form'

import { uploadImage as UPLOAD_IMAGE } from '@gql'
import { usePropertyOnboarding } from '@usePropertyOnboarding'

import { caretakerSchema } from '../validations'
import { type CaretakerForm, type OnboardingStep } from '../types'

function Caretaker (): React.Node {
  const [uploadImage, { loading: uploadingImage }] = useMutation(UPLOAD_IMAGE)
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { setStep, caretakerForm, setCaretakerForm } = usePropertyOnboarding()
  const { register, handleSubmit, setValue, formState: { errors }, watch } = useForm<CaretakerForm>({
    defaultValues: caretakerForm,
    resolver: yupResolver(caretakerSchema)
  })
  const { getRootProps, getInputProps } = useDropzone({
    accept: {
      'image/*': ['.jpeg', '.jpg', '.png', '.gif']
    },
    multiple: false,
    disabled: uploadingImage,
    onDrop: async (acceptedFiles) => {
      // TODO upload to s3
      const res = await uploadImage({
        variables: {
          file: acceptedFiles[0]
        }
      })
      // @ts-expect-error
      setValue('caretakerForm.idVerification', res?.data.uploadImage)
    }
  })

  const idImg = watch('caretaker.idVerification')
  // TODO caretaker phone verification flow
  const onSubmit: SubmitHandler<CaretakerForm> = data => {
    setCaretakerForm(data)
    // TODO Send verification code to phone
    onOpen()
    // TODO proceed
  }
  const goBack = (): Dispatch<SetStateAction<OnboardingStep>> => {
    setStep('pricing')
  }

  return (
    <form onSubmit={() => handleSubmit(onSubmit)}>
      <Stack align="center" direction={{ base: 'column', md: 'row' }} spacing={{ base: 4, md: 6 }}>
        <Modal isCentered isOpen={isOpen} onClose={onClose}>
          <ModalContent>
            <ModalHeader>Verify Phone</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              Phone Verification
            </ModalBody>
          </ModalContent>
        </Modal>
        <Box w="100%">
          <FormControl isInvalid={Boolean(errors?.firstName)}>
            <FormLabel>First Name</FormLabel>
            <Input
              {...register('firstName')}
            />
            {(errors.firstName != null) && <FormErrorMessage>{errors?.firstName.message}</FormErrorMessage>}
          </FormControl>
          <FormControl isInvalid={Boolean(errors?.lastName)}>
            <FormLabel>Last Name</FormLabel>
            <Input
              {...register('lastName')}
            />
            {((errors?.lastName) != null) && <FormErrorMessage>{errors?.lastName.message}</FormErrorMessage>}
          </FormControl>
          <FormControl isInvalid={Boolean(errors?.phoneNumber)}>
            <FormLabel>Phone Number</FormLabel>
            <Input
              {...register('phoneNumber')}
              type="number"
            />
            {((errors?.phoneNumber) != null) && <FormErrorMessage>{errors?.phoneNumber.message}</FormErrorMessage>}
          </FormControl>
        </Box>
        <FormControl isInvalid={Boolean(errors?.idVerification)}>
          <Textarea
            as={Center}
            {...getRootProps({ className: 'dropzone' })}
            p={4}
            minH={{ base: '80px', md: '100px' }}
            cursor="pointer"
            borderRadius="md"
            border="2px dashed"
            borderColor="chakra-border-color"
            spacing={4}
          >
            {idImg && <Image
              src={idImg}
              loading="eager"
            />}
            {!uploadingImage && <Icon as={FaUpload} />}
            {uploadingImage && <Spinner size="lg" />}
          </Textarea>
          <input {...register('idVerification')} {...getInputProps()} />
          {((errors?.idVerification) != null) && <FormErrorMessage>{errors?.idVerification.message}</FormErrorMessage>}
          <FormHelperText>Government issued document</FormHelperText>
        </FormControl>
      </Stack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button colorScheme="green" leftIcon={<ArrowBackIcon />} onClick={goBack}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Caretaker
