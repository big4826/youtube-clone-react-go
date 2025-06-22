import { Controller, useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Input, Button, Form } from "antd";

// Zod schema
const schema = z.object({
  email: z.string().email({ message: "Invalid email address" }),
  password: z.string().min(6, { message: "Minimum 6 characters required" }),
});

// Infer form data type from schema
type LoginFormData = z.infer<typeof schema>;

export const PageLogin = () => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(schema),
  });

  const onSubmit = (data: LoginFormData) => {
    console.log("Submitted data:", data);
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* <Flex gap="middle" vertical justify="center" align="center"> */}
        <Form.Item
          //   label="Email"
          validateStatus={errors.email ? "error" : ""}
          help={errors.email?.message}
          style={{ width: "100%" }}
        >
          <Controller
            name="email"
            control={control}
            render={({ field }) => <Input {...field} />}
          />
        </Form.Item>

        <Form.Item
          //   label="Password"
          validateStatus={errors.password ? "error" : ""}
          help={errors.password?.message}
          style={{ width: "100%" }}
        >
          <Controller
            name="password"
            control={control}
            render={({ field }) => <Input.Password {...field} />}
          />
        </Form.Item>
      {/* </Flex> */}
      <Form.Item>
        <Button type="primary" htmlType="submit" block>
          Login
        </Button>
      </Form.Item>
    </form>
  );
};
