import { useMutation } from '@apollo/client'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Box, Center, Button, HStack, Image, FormControl, FormErrorMessage, FormHelperText, FormLabel, Icon, Input, Spacer, Stack, Textarea, useDisclosure, Spinner, Select } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { useDropzone } from 'react-dropzone'
import { useForm, type SubmitHandler } from 'react-hook-form'
import { FaUpload } from 'react-icons/fa'

import { VerificationModal } from '../components'
import { type CaretakerForm } from '../types'
import { caretakerSchema } from '../validations'

import { uploadImage as UPLOAD_IMAGE, sendVerificationCode as SEND_VERIFICATION_CODE, } from '@gql'
import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Caretaker = (): JSX.Element => {
  const [uploadImage, { loading: uploadingImage }] = useMutation(UPLOAD_IMAGE)
  const [sendVerification, { loading: sendingVerification }] = useMutation(SEND_VERIFICATION_CODE)
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { setStep, caretakerForm, setCaretakerForm, caretakerVerified } = usePropertyOnboarding()
  const { register, handleSubmit, setValue, formState: { errors }, trigger, watch } = useForm<CaretakerForm>({
    defaultValues: caretakerForm,
    resolver: yupResolver(caretakerSchema)
  })
  const handleDrop = async (acceptedFiles: File[]) => {
    const res = await uploadImage({
      variables: {
        file: acceptedFiles[0]
      }
    })
    setValue('idVerification', res?.data.uploadImage)
    trigger("idVerification")
  }
  const { getRootProps, getInputProps } = useDropzone({
    accept: {
      'image/*': ['.jpeg', '.jpg', '.png', '.gif']
    },
    multiple: false,
    disabled: uploadingImage,
    onDrop: handleDrop,
  })

  // Watch verification img changes
  const idImg = watch('idVerification')

  // Start caretaker verification flow
  const onSubmit: SubmitHandler<CaretakerForm> = async data => {
    setCaretakerForm(data)
    // Send verification code to phone
    if (!caretakerVerified || (data.phoneNumber != caretakerForm.phoneNumber)) {
      await sendVerification({
        variables: {
          input: {
            phone: `${caretakerForm.countryCode}${data.phoneNumber}`,
            countryCode: "KE",
          },
        },
        // Proceed to next step once successfull
        onCompleted: data => {
          const status = data?.sendVerificationCode.success
          if (status === "pending") {
            onOpen()
          }
        },
      })
    }
  }
   
  const goBack = (): void => {
    setStep('pricing')
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack align="center" direction={{ base: 'column', md: 'row' }} spacing={{ base: 4, md: 6 }}>
        <VerificationModal
          onClose={onClose}
          isOpen={isOpen}
        />
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
          <FormControl isInvalid={Boolean(errors?.phoneNumber || errors?.countryCode)}>
            <FormLabel>Phone Number</FormLabel>
            <HStack>
              <Select {...register("countryCode")}>
                <option value="+254">+254</option>
              </Select>
              <Input
                {...register('phoneNumber')}
                type="number"
              />
            </HStack>
            {((errors?.phoneNumber) != null) && <FormErrorMessage>{errors?.phoneNumber.message}</FormErrorMessage>}
          </FormControl>
        </Box>
        <FormControl isInvalid={Boolean(errors?.idVerification)}>
        <FormLabel> Identification Document</FormLabel>
          <Textarea
            as={Center}
            {...getRootProps({ className: 'dropzone' })}
            p={4}
            minH={{ base: '80px', md: '100px' }}
            cursor="pointer"
            h="auto"
            justify={idImg ? "start" : "center"}
            borderRadius="md"
            border="2px dashed"
            borderColor="chakra-border-color"
            spacing={4}
          >
            {idImg && !uploadingImage && <Image
              src={idImg}
              loading="eager"
              maxW={{
                base: "100px",
                md: "200px"
              }}
              alt="ID Verification"
            />}
            {!idImg && !uploadingImage && <Icon as={FaUpload} />}
            {uploadingImage && <Spinner size="lg" />}
          </Textarea>
          <input {...register('idVerification')} {...getInputProps()} />
          {((errors?.idVerification) != null) && <FormErrorMessage>{errors?.idVerification.message}</FormErrorMessage>}
          <FormHelperText>Government issued document</FormHelperText>
        </FormControl>
      </Stack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button colorScheme="green" disabled={sendingVerification} leftIcon={<ArrowBackIcon />} onClick={goBack}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" disabled={sendingVerification} isLoading={sendingVerification} rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Caretaker
