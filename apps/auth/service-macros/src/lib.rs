extern crate proc_macro;
use convert_case::{Case, Casing};
use proc_macro2::{Span, TokenStream};
use quote::quote;
// use service_core::Repository;
use syn::{parse::Parse, parse_macro_input, Ident, Token, Visibility};

struct Input {
    pub service_name: Ident,
    pub dependencies: Vec<Ident>,
    pub visibility: Visibility,
}

impl Parse for Input {
    fn parse(input: syn::parse::ParseStream) -> syn::Result<Self> {
        let visibility = input.parse::<Visibility>()?;
        let service_name = input.parse::<Ident>()?;
        input.parse::<Token![use]>()?;
        let mut dependencies: Vec<Ident> = Vec::new();

        while !input.is_empty() && input.peek(Ident) {
            let ty = input.parse::<Ident>()?;
            dependencies.push(ty);
            let _ = input.parse::<Token![,]>();
        }

        Ok(Input {
            service_name,
            dependencies,
            visibility,
        })
    }
}

#[proc_macro]
pub fn service(input: proc_macro::TokenStream) -> proc_macro::TokenStream {
    let input = parse_macro_input!(input as Input);

    let struct_stream = get_struct(&input);
    let impl_stream = get_impl(&input);

    let expanded = quote! {
        #struct_stream

        #impl_stream
    };

    // println!("{}", expanded);
    proc_macro::TokenStream::from(expanded)
}

fn get_struct(input: &Input) -> TokenStream {
    let Input {
        service_name,
        dependencies,
        visibility,
    } = input;

    let attrs = dependencies.iter().map(|x| {
        let mut attr_name = x.to_string().to_case(Case::Snake);
        attr_name.push_str("_repository");
        let attr_ident = Ident::new(&attr_name, Span::call_site());

        quote! {
            pub #attr_ident: service_core::Repository<#x>,
        }
    });

    quote! {
        #visibility struct #service_name {
            #(#attrs)*
        }
    }
}

fn get_impl(input: &Input) -> TokenStream {
    let Input {
        service_name,
        dependencies,
        ..
    } = input;

    let fn_args = dependencies.iter().map(|x| {
        let mut attr_name = x.to_string().to_case(Case::Snake);
        attr_name.push_str("_repository");
        let attr_ident = Ident::new(&attr_name, Span::call_site());

        quote! {
            #attr_ident: service_core::Repository<#x>,
        }
    });

    let attr_body = dependencies.iter().map(|x| {
        let mut attr_name = x.to_string().to_case(Case::Snake);
        attr_name.push_str("_repository");
        let attr_ident = Ident::new(&attr_name, Span::call_site());

        quote! {
            #attr_ident,
        }
    });

    quote! {
        impl #service_name {
            pub fn new(#(#fn_args)*) -> Self {
                Self {
                    #(#attr_body)*
                }
            }
        }
    }
}
